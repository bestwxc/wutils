package main

import (
	"fmt"
	"github.com/bestwxc/wutils"
)

func main() {
	src1 := "test/image/pic1.png"
	src2 := "test/image/pic2.jpg"
	fmt.Println(wutils.PathIsFileType(wutils.PngFile, src1))
	fmt.Println(wutils.PathIsFileType(wutils.PngFile, src2))
	fmt.Println(wutils.PathIsFileType(wutils.JpegFile, src1))
	fmt.Println(wutils.PathIsFileType(wutils.JpegFile, src2))
}
