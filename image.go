package wutils

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	DefaultFileMode     = 0766
	CopyWhenUnsupported = 1
)

// CompressImage 压缩图片
func CompressImage(sourcePath string, targetPath string, format string, quality int) error {
	return CompressImageExt(sourcePath, targetPath, format, quality, DefaultFileMode, CopyWhenUnsupported)
}

// CompressImageExt 压缩图片的扩展方法
func CompressImageExt(sourcePath string, targetPath string, format string, quality int, perm os.FileMode, unsupported int) error {
	if len(format) > 0 && !SupportFormat(format) {
		panic("目标格式不支持,format:" + format)
	}
	// 打开文件
	imgfile, err1 := os.Open(sourcePath)
	defer imgfile.Close()
	if err1 != nil {
		return err1
	}
	// 解码
	img, sourceFormat, err2 := image.Decode(imgfile)
	if err2 != nil {
		return err2
	}
	if !SupportFormat(sourceFormat) {
		if unsupported == 1 {
			// 复制，否则跳过
			CopyFile(sourcePath, targetPath)
		}
		return nil
	}
	// 创建输出文件
	newfile, err3 := os.Create(targetPath)
	defer newfile.Close()
	if err3 != nil {
		return err3
	}
	targetFormat := sourceFormat
	if len(format) > 0 {
		targetFormat = format
	}
	var err4 error
	switch targetFormat {
	case "jpeg":
		err4 = jpeg.Encode(newfile, img, &jpeg.Options{Quality: quality})
	case "png":
		err4 = png.Encode(newfile, img)
	case "gif":
		err4 = gif.Encode(newfile, img, &gif.Options{})
	default:
		panic("不支持的图片类型")
	}
	if err4 != nil {
		return err4
	}
	return nil
}

// SupportFormat 是否支持图片格式
func SupportFormat(format string) bool {
	return format == "jpeg" || format == "png" || format == "gif"
}
