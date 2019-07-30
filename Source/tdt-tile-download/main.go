package main

import (
	"IrisTestProject/Source/tdt-tile-download/TDT"
	"IrisTestProject/Source/tdt-tile-download/parseIni"
	"IrisTestProject/Source/tdt-tile-download/util"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main()  {
	configFile:=util.GetCurrentPath()+"\\config.ini"
	isExist,err:=util.FileExists(configFile)
	if err!=nil{
		fmt.Printf("配置文件检测失败！错误信息为：%s\n", err)
		return
	}else {
		if !isExist{
			fmt.Printf("配置文件不存在！\n")
			return
		}
	}

	c,err := parseIni.ReadIniFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxConns,err := c.GetConfigToInt("maxConns")
	if err!=nil{
		fmt.Printf("配置文件中maxConns检测失败！错误信息为：%s\n", err)
		return
	}
	sleep,err := c.GetConfigToInt("sleep")
	if err!=nil{
		fmt.Printf("配置文件中sleep检测失败！错误信息为：%s\n", err)
		return
	}
	overwrite,err := c.GetConfigToBool("overwrite")
	if err!=nil{
		fmt.Printf("配置文件中overwrite检测失败！错误信息为：%s\n", err)
		return
	}
	projection,err := c.GetConfigToString("projection")
	if err!=nil{
		fmt.Printf("配置文件中projection检测失败！错误信息为：%s\n", err)
		return
	}
	serviceMode,err := c.GetConfigToString("serviceMode")
	if err!=nil{
		fmt.Printf("配置文件中serviceMode检测失败！错误信息为：%s\n", err)
		return
	}
	layer,err := c.GetConfigToString("layer")
	if err!=nil{
		fmt.Printf("配置文件中layer检测失败！错误信息为：%s\n", err)
		return
	}
	tileType,err := c.GetConfigToString("tileType")
	if err!=nil{
		fmt.Printf("配置文件中tileType检测失败！错误信息为：%s\n", err)
		return
	}
	outPath,err := c.GetConfigToString("outPath")
	if err!=nil{
		fmt.Printf("配置文件中outPath检测失败！错误信息为：%s\n", err)
		return
	}
	minLevel,err := c.GetConfigToInt("level.minLevel")
	if err!=nil{
		fmt.Printf("配置文件中minLevel检测失败！错误信息为：%s\n", err)
		return
	}
	maxLevel,err := c.GetConfigToInt("level.maxLevel")
	if err!=nil{
		fmt.Printf("配置文件中maxLevel检测失败！错误信息为：%s\n", err)
		return
	}
	minX,err := c.GetConfigToFloat64("bbox.minX")
	if err!=nil{
		fmt.Printf("配置文件中minX检测失败！错误信息为：%s\n", err)
		return
	}
	minY,err := c.GetConfigToFloat64("bbox.minY")
	if err!=nil{
		fmt.Printf("配置文件中minY检测失败！错误信息为：%s\n", err)
		return
	}
	maxX,err := c.GetConfigToFloat64("bbox.maxX")
	if err!=nil{
		fmt.Printf("配置文件中maxX检测失败！错误信息为：%s\n", err)
		return
	}
	maxY,err := c.GetConfigToFloat64("bbox.maxY")
	if err!=nil{
		fmt.Printf("配置文件中maxY检测失败！错误信息为：%s\n", err)
		return
	}

	fmt.Printf("  配置信息如下：\n")
	fmt.Printf("	maxConns=%v\n",maxConns)
	fmt.Printf("	sleep=%v\n",sleep)
	fmt.Printf("	overwrite=%v\n",overwrite)
	fmt.Printf("	projection=%v\n",projection)
	fmt.Printf("	projection=%v\n",serviceMode)
	fmt.Printf("	layer=%v\n",layer)
	fmt.Printf("	tileType=%v\n",tileType)
	fmt.Printf("	outPath=%v\n",outPath)
	fmt.Printf("	level.minLevel=%v\n",minLevel)
	fmt.Printf("	level.maxLevel=%v\n",maxLevel)
	fmt.Printf("	bbox.minX=%v\n",minX)
	fmt.Printf("	bbox.minY=%v\n",minY)
	fmt.Printf("	bbox.maxX=%v\n",maxX)
	fmt.Printf("	bbox.maxY=%v\n",maxY)

	fmt.Printf("  是否继续(y/n)：\n")

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command :=strings.ToLower(string(data))

	if command=="y"{
		http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = maxConns
		http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost=maxConns
		http.DefaultTransport.(*http.Transport).MaxIdleConns=maxConns

		TDT.DownloadMulTiles(maxConns,sleep,overwrite,projection,serviceMode,layer,tileType,outPath,minLevel,maxLevel,minX,minY,maxX,maxY)
	}else if command=="n"{
		return
	}
}