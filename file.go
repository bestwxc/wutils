package wutils

import (
	"fmt"
	"io"
	"os"
)

// CopyFileBuffSize 复制用的buff大小 4k
const CopyFileBuffSize = 1024 * 4

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
	return CopyFileWithBuff(source, destination, CopyFileBuffSize)
}

// CopyFileWithBuff 使用buff讲一个打开的文件复制到另一个打开的文件
func CopyFileWithBuff(srcFile *os.File, dstFile *os.File, bufSize int) (int64, error) {
	var size int64 = 0
	buf := make([]byte, bufSize)
	for {
		n, err2 := srcFile.Read(buf)
		if err2 != nil && err2 != io.EOF {
			return -1, err2
		}
		if n == 0 {
			break
		}
		if _, err3 := dstFile.Write(buf[:n]); err3 != nil {
			return -1, err3
		}
		size += int64(n)
	}
	return size, nil
}
