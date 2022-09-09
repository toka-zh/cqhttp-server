package pkg

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func GetRandFileAbsPath(path string) string {
	file := GetRandFile(path)
	if file == "" {
		return ""
	}
	return "file:///" + file
}

// GetRandFile 随机获取目录下的文件绝对路径
// todo 做一个资源池
//
//	获取目录下的文件(目录可配置),然后随机获取一个文件的绝对路径
func GetRandFile(path string) string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return ""
	}

	randInt := RandInt(len(dir))
	abs, _ := filepath.Abs(path + "/" + dir[randInt].Name())
	return abs
}
