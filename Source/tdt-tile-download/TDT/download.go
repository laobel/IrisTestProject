package TDT

import (
	"IrisTestProject/Source/tdt-tile-download/ellipsoid"
	"IrisTestProject/Source/tdt-tile-download/users"
	"IrisTestProject/Source/tdt-tile-download/webMercatorProjection"
	"IrisTestProject/Source/tdt-tile-download/webMercatorTilingScheme"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func downloadTile(x,y,z int,url,outDir,outFileType string,overwrite bool,currentNum *int,waitGroup *sync.WaitGroup){
	outPath:=fmt.Sprintf(outDir+"/%d/%d", z,y)
	os.MkdirAll(outPath, 0666)

	outfile:=fmt.Sprintf(outPath+"/%d.%s",x,outFileType)

	if !overwrite{//不替换已有瓦片
		_, err := os.Stat(outfile)
		if err == nil {
			//log.Println("文件已经存在")
			defer waitGroup.Done()
			*currentNum--
			return
		}
	}

	client := &http.Client{}

	url=fmt.Sprintf(url,subdomains[rand.Intn(8)],x,y,z,users.RandomToken())
	req, err := http.NewRequest("POST", url,nil)
	if err != nil {
		log.Println(err)
		defer waitGroup.Done()
		*currentNum--
		return
	}

	req.Header.Set("User-Agent", users.RandomAgent())

	resp, err := client.Do(req)
	if err != nil {
		*currentNum++
		time.Sleep(1 * time.Second)//休眠10秒
		waitGroup.Add(1)
		downloadTile(x,y,z,url,outDir,outFileType,overwrite,currentNum,waitGroup)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		defer waitGroup.Done()
		*currentNum--
		return
	}

	ioutil.WriteFile(outfile, body, 0644)

	//计数器-1
	defer waitGroup.Done()
	*currentNum--
}

func DownloadMulTiles(maxConns,sleep int,overwrite bool,projection,serviceMode,layer,tileType,outPath string,minLevel,maxLevel int,	minX,minY,maxX,maxY float64)  {
	fmt.Printf("========================================================================================\n")

	var waitGroup=new(sync.WaitGroup)
	var ellipsoid=ellipsoid.WGS84Ellipsoid()
	var webMercatorProjection=webMercatorProjection.NewWWebMercatorProjection(*ellipsoid)
	var webMercatorTilingScheme=webMercatorTilingScheme.NewWebMercatorTilingScheme(*ellipsoid,*webMercatorProjection,1,1)

	var url=Url(projection,layer,serviceMode)

	var currentNum=0
	var num float64=0

	var total float64=0
	for  z:=minLevel;z<=maxLevel;z++ {
		xTitleMin, yTitleMin := webMercatorTilingScheme.LonLatToTileXY(minX, maxY, z)
		xTitleMax, yTitleMax := webMercatorTilingScheme.LonLatToTileXY(maxX, minY, z)
		total+=float64((xTitleMax-xTitleMin+1)*(yTitleMax-yTitleMin+1))
	}

	var flag=0

	now := time.Now()
	fmt.Printf("开始执行。当前时间为：%v\n",time.Now())

	for  z:=minLevel;z<=maxLevel;z++ {
		xTitleMin,yTitleMin:=webMercatorTilingScheme.LonLatToTileXY(minX,maxY,z)
		xTitleMax,yTitleMax:=webMercatorTilingScheme.LonLatToTileXY(maxX,minY,z)

		for x:=xTitleMin;x<=xTitleMax;x++{
			for y:=yTitleMin;y<=yTitleMax;y++{
				currentNum++
				for currentNum>=maxConns-1{
					time.Sleep(time.Duration(sleep) * time.Second)//休眠 秒
				}

				if int(int(num/total*100)/10)*10==flag{
					fmt.Printf("已完成：%d%s,耗时:%v\n",flag,"%",time.Now().Sub(now))
					flag+=10
				}
				waitGroup.Add(1)
				num++
				go downloadTile(x,y,z,url,outPath,tileType,overwrite,&currentNum,waitGroup)
			}
		}
	}
	//等待所有协程操作完成
	waitGroup.Wait()
	fmt.Printf("已完成：%d%s,耗时:%v\n",flag,"%",time.Now().Sub(now))
	fmt.Printf("全部处理完成！下载总时间:%v\n", time.Now().Sub(now))
}