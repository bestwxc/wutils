package main

import (
	"fmt"
	"github.com/bestwxc/wutils"
	"os"
)

func main() {
	TestImage()
}

func TestImage() {
	src1 := "test/image/pic1.png"
	dst1 := "test/image/pic11.png"
	dst12 := "test/image/pic12.jpg"
	dst13 := "test/image/pic13.png"

	src2 := "test/image/pic2.jpg"
	dst2 := "test/image/pic21.jpg"
	dst22 := "test/image/pic22.jpg"
	dst23 := "test/image/pic23.png"
	if wutils.FileExists(dst1) {
		os.Remove(dst1)
	}
	if wutils.FileExists(dst12) {
		os.Remove(dst12)
	}
	if wutils.FileExists(dst13) {
		os.Remove(dst13)
	}
	if wutils.FileExists(dst2) {
		os.Remove(dst2)
	}
	if wutils.FileExists(dst22) {
		os.Remove(dst22)
	}
	if wutils.FileExists(dst23) {
		os.Remove(dst23)
	}
	err1 := wutils.CompressImage(src1, dst1, "", 50)
	fmt.Println(err1)
	err2 := wutils.CompressImage(src2, dst2, "", 50)
	fmt.Println(err2)

	err21 := wutils.CompressImage(src1, dst12, "jpeg", 50)
	fmt.Println(err21)
	err22 := wutils.CompressImage(src2, dst22, "jpeg", 50)
	fmt.Println(err22)

	err31 := wutils.CompressImage(src1, dst13, "png", 50)
	fmt.Println(err31)
	err32 := wutils.CompressImage(src2, dst23, "png", 50)
	fmt.Println(err32)
}
