package wutils

import (
	"fmt"
	"io"
	"os"
)

const CopyFileBuffSize = 1024

// CopyFile 复制文件
func CopyFile(src string, dst string) (int64, error) {
	if !FileExists(src) {
		return -1, fmt.Errorf("源文件不存在")
	}
	source, err := os.Open(src)
	if err != nil {
		return -1, err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return -1, err
	}
	defer destination.Close()
	buf := make([]byte, CopyFileBuffSize)
	var size int64 = 0
	for {
		n, err2 := source.Read(buf)
		if err2 != nil && err2 != io.EOF {
			return -1, err2
		}
		if n == 0 {
			break
		}
		if _, err3 := destination.Write(buf[:n]); err != nil {
			return -1, err3
		}
		size += int64(n)
	}
	return size, nil
}
