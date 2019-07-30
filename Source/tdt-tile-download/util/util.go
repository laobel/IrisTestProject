package util

import (
	"math"
	"os"
	"os/exec"
	"path/filepath"
)

//度转弧度
func Radians(deg float64)float64 {
	return deg * (math.Pi / 180.0)
}

/*获取当前文件执行的路径*/
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileExists(filepath string)(bool,error)  {
	fInfo, err := os.Stat(filepath)
	if err == nil {
		if fInfo.IsDir() {
			return false,nil
		} else {
			return  true,nil
		}
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

