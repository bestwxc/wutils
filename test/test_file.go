package main

import (
	"fmt"
	"github.com/bestwxc/wutils"
	"os"
)

func main() {
	TestFile()
}

func TestFile() {
	src := "test/file/t1.txt"
	dst := "test/file/t2.txt"
	if wutils.FileExists(dst) {
		os.Remove(dst)
	}
	_, err := wutils.CopyFile(src, dst)
	msg := "success"
	if err != nil {
		msg = "复制出现错误,error:" + err.Error()
	}
	if !wutils.FileExists(dst) {
		msg = "复制出现错误，目标文件不存在"
	}
	fmt.Println("test FileExists result: ", msg)
}
