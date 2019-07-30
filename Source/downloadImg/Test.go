//使用golang中sync.WaitGroup来实现协程同步

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var waitGroup=new(sync.WaitGroup)

func download(i int ){
	url := fmt.Sprintf("http://services.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/8/113/%d.jpg", i)
	fmt.Printf("开始下载:%s\n", url)
	res,err := http.Get(url)
	if err != nil || res.StatusCode != 200{
		fmt.Printf("下载失败:%s", res.Request.URL)
	}
	fmt.Printf("开始读取文件内容,url=%s\n", url)
	data ,err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Printf("读取数据失败")
	}
	os.MkdirAll("D:/GoPath/src/IrisTestProject/arcgisImg/tile/8/113", 0666)
	ioutil.WriteFile(fmt.Sprintf("D:/GoPath/src/IrisTestProject/arcgisImg/tile/8/113/%d.jpg", i), data, 0644)

	//if failed, sudo chmod 777 pic2016/

	//计数器-1
	waitGroup.Done()
}


var agents =[20] string {
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
	"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/532.9 (KHTML, like Gecko) Chrome/5.0.310.0 Safari/532.9",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Chrome/7.0.514.0 Safari/534.7",
	"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/534.14 (KHTML, like Gecko) Chrome/9.0.601.0 Safari/534.14",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.14 (KHTML, like Gecko) Chrome/10.0.601.0 Safari/534.14",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.20 (KHTML, like Gecko) Chrome/11.0.672.2 Safari/534.20",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.27 (KHTML, like Gecko) Chrome/12.0.712.0 Safari/534.27",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.24 Safari/535.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/4.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/5.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/9.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/10.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/11.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.3683.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17134",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"}

var tokens=[20] string{
	"1072d95046f18e67463ce40d645a9b8d",
	"a6484e08c7c9e7fd9bec784163b9ef18",
	"b2b8f791570d96f6cf46898747e7ce1a",
	"9aebc7723eee1ef07a73f5ddfc0c1df4",
	"eac0051c446319b99c65da060f3e83e4",
	"86c9e10dfe14ced9f43901b6f4c9e983",
	"85b88ce10c15f390ee75bf571688b3b7",
	"03a74352f8945d4e011b1914e0527514",
	"9024bb5d2e154746bb513878231cc0cf",
	"c1d6b49adb2ba817109873dbc13becb4",
	"28b495e4df789d971d2ae77b01a55a55",
	"997487c2aa6dc93d84169f293ae2073d",
	"1dfcf1a2b70604242eb5c6abe8ee9703",
	"19a3cd1b08ddbfeb3701f58b4843bac0",
	"9ce21c72792f415f5ce0b96e29e00d52",
	"b0ad70dc306d789204ddb4ec0b7c2b4d",
	"854a6c21a3f15f27dab5b2d676ad2321",
	"037732e0f605c2b98884616a4584d38f",
	"b976a2cd81f5fab373ced07d17b9aa81",
	"a4cccd543a7a9fbe8b85f04746eb2753"}

var  subdomains=[8] string{"t0","t1","t2","t3","t4","t5","t6","t7"}

var resolution=map[int]float64{
	1: 0.70312500015485435,
	2: 0.35156250007742718,
	3: 0.17578125003871359,
	4:0.0878906250193568,
	5:0.0439453125096784,
	6:0.0219726562548392,
	7:0.0109863281274196,
	8:0.0054931640637098,
	9:0.0027465820318549957,
	10:0.0013732910159274978,
	11:0.00068664549607834132,
	12:0.00034332275992416907,
	13:0.00017166136807812298,
	14:8.5830684039061379E-05,
	15:4.2915342019530649E-05,
	16:2.1457682893727977E-05,
	17:1.0728841446864E-05,
	18:5.3644207234319882E-06,
	19:2.6822103617159941E-06,
	20:1.341105180858E-06}



//度转弧度
func radians(deg float64)float64{
	return deg * (math.Pi / 180.0)
}

func getNumberOfXTilesAtLevel(numberOfLevelZeroTilesX,level uint) float64{
	return float64(numberOfLevelZeroTilesX << level)
}
func getNumberOfYTilesAtLevel(numberOfLevelZeroTilesY,level uint) float64{
	return float64(numberOfLevelZeroTilesY << level)
}

var maximumRadius=math.Max(6378137.0, 6356752.3142451793)
var semimajorAxisTimesPi = maximumRadius * math.Pi
type rectangleInMeters struct {
	x float64
	y float64
}

var MaximumLatitude=math.Pi/2.0 - (2.0 * math.Atan(math.Exp(-math.Pi)))

func geodeticLatitudeToMercatorAngle(lat float64) float64{
	if lat > MaximumLatitude {
		lat = MaximumLatitude
	} else if lat < -MaximumLatitude {
		lat = -MaximumLatitude
	}
	var sinLatitude = math.Sin(lat)

	return 0.5 * math.Log((1.0 + sinLatitude) / (1.0 - sinLatitude))
}

func project (lon,lat float64) (float64,float64) {
	longitude:=radians(lon)
	latitude:=radians(lat)

	var semimajorAxis = maximumRadius

	var x = longitude * semimajorAxis
	var y = geodeticLatitudeToMercatorAngle(latitude) * semimajorAxis

	return x,y
}


func deg2num1(numberOfLevelZeroTilesX,numberOfLevelZeroTilesY uint,rectangleNortheastInMeters ,rectangleSouthwestInMeters rectangleInMeters,lon,lat float64,zoom int)(int,int){
	var xTiles = getNumberOfXTilesAtLevel(numberOfLevelZeroTilesX,uint(zoom))
	var yTiles = getNumberOfYTilesAtLevel(numberOfLevelZeroTilesY,uint(zoom))

	var overallWidth = rectangleNortheastInMeters.x - rectangleSouthwestInMeters.x
	var xTileWidth = overallWidth / xTiles
	var overallHeight = rectangleNortheastInMeters.y - rectangleSouthwestInMeters.y
	var yTileHeight = overallHeight / yTiles

	webMercatorPositionX, webMercatorPositionY:= project(lon,lat)
	var distanceFromWest = webMercatorPositionX - rectangleSouthwestInMeters.x
	var distanceFromNorth = rectangleNortheastInMeters.y - webMercatorPositionY

	var xTileCoordinate = int(distanceFromWest / xTileWidth) | 0.0
	if xTileCoordinate >= int(xTiles) {
		xTileCoordinate = int(xTiles) - 1
	}
	var yTileCoordinate = int(distanceFromNorth / yTileHeight) | 0
	if yTileCoordinate >= int(yTiles) {
		yTileCoordinate = int(yTiles) - 1
	}

	return xTileCoordinate,yTileCoordinate
}


//经纬度反算切片行列号 3857坐标系
//x = ceiling（（180 + X）/s*256）-1

//y = ceiling（（90-Y）/s*256）-1；
func deg2num(lon,lat float64,zoom int)(int,int){
	s:=resolution[zoom]

	x:=int(math.Ceil((180.0+lon)/(s*256.0))-1)
	y:=int(math.Ceil((90.0-lat)/(s*256.0))-1)

	if x<0 {
		x = 0
	}
	if y<0{
		y=0
	}

	return x,y
}


//下载天地图数据
func downloadTDT(x,y,z int,url string,outDir string,overwrite bool,num *int){
	outPath:=fmt.Sprintf(outDir+"/%d/%d", z,y)
	os.MkdirAll(outPath, 0666)

	outfile:=fmt.Sprintf(outPath+"/%d.png",x)

	if !overwrite{//不替换已有瓦片
		_, err := os.Stat(outfile)
		if err == nil {
			//log.Println("文件已经存在")
			waitGroup.Done()
			*num--
			return
		}
	}

	client := &http.Client{

	}

	url=fmt.Sprintf(url,subdomains[rand.Intn(8)],x,y,z,tokens[rand.Intn(20)])
	req, err := http.NewRequest("POST", url,nil)
	if err != nil {
		log.Println(err)
		defer waitGroup.Done()
		*num--
		return
	}

	req.Header.Set("User-Agent", agents[rand.Intn(20)])

	resp,err := client.Do(req)
	if err!=nil {
		fmt.Printf("错误为:%v\n", err)
		*num++
		waitGroup.Add(1)
		time.Sleep(1 * time.Second)//休眠10秒
		go downloadTDT(x,y,z,url,outDir,false,num)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
	fmt.Printf("%v大小为:%v\n", url,len(body))
	if err != nil {
		log.Println(err)
		defer waitGroup.Done()
		*num--
		return
	}

	ioutil.WriteFile(outfile, body, 0644)

	//计数器-1
	defer waitGroup.Done()
	*num--
}

func main()  {
	//创建多个协程，同时下载多个图片
	now := time.Now()

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 500
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost=500
	http.DefaultTransport.(*http.Transport).MaxIdleConns=500

	var outPath="D:/apache-tomcat-9.0.14/webapps/Download_TDT/卫星影像_TDT_W"
	var url="http://%s.tianditu.gov.cn/cia_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=cia&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILECOL=%d&TILEROW=%d&TILEMATRIX=%d&tk=%s"

	/*zoomMin,zoomMax:=1,6
	xMin,xMax:=-180.0,180.0
	yMin,yMax:=-90.0,90.0*/

	zoomMin,zoomMax:=1,18
	xMin,xMax:=120.784717852732,122.310761982605
	yMin,yMax:=30.604597748298,31.904507540401


	var rectangleNortheastInMeters rectangleInMeters
	rectangleNortheastInMeters.x=semimajorAxisTimesPi
	rectangleNortheastInMeters.y=semimajorAxisTimesPi

	var rectangleSouthwestInMeters rectangleInMeters
	rectangleSouthwestInMeters.x=-semimajorAxisTimesPi
	rectangleSouthwestInMeters.y=-semimajorAxisTimesPi

	var maxIntance=1000

	var n int =0
	for  z:=zoomMin;z<=zoomMax;z++ {
		xTitleMin,yTitleMin:=deg2num1(1,1,rectangleNortheastInMeters,rectangleSouthwestInMeters,xMin,yMax,z)
		xTitleMax,yTitleMax:=deg2num1(1,1,rectangleNortheastInMeters,rectangleSouthwestInMeters,xMax,yMin,z)

		for x:=xTitleMin;x<=xTitleMax;x++{
			for y:=yTitleMin;y<=yTitleMax;y++{
				waitGroup.Add(1)
				n=n+1
				/*for sleep{
					time.Sleep(5 * time.Second)//休眠10秒
				}*/
				for n>=maxIntance-1{
					time.Sleep(1 * time.Second)//休眠10秒
				}

				go downloadTDT(x,y,z,url,outPath,false,&n)

				//fmt.Printf("%d\r\n",n)
			}
		}
	}
/*
	for i :=100; i<300; i++ {
		//计数器+1
		waitGroup.Add(1)
		go download(i)
	}*/

	//等待所有协程操作完成
	waitGroup.Wait()
	fmt.Printf("下载总时间:%v\n", time.Now().Sub(now))
}