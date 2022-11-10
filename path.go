package wutils

import (
	"os"
	"path"
)

// FileExists 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// NotDir 判断所给路径是否为文件
func NotDir(path string) bool {
	return !IsDir(path)
}

// CreateFolderIfNotExists 创建临时文件路径，如果他不存在的话
func CreateFolderIfNotExists(filePath string, perm os.FileMode) {
	folder := path.Dir(filePath)
	os.MkdirAll(folder, perm)
}
