package wutils

import "os"

// FileType 文件类型
type FileType struct {
	Code        string
	Name        string
	Suffix      []string
	MagicNumber [][]byte
}

// PathIsFileType 判断指定路径的文件是否与文件类型匹配
func PathIsFileType(ft *FileType, p string) bool {
	f, err := os.Open(p)
	if err != nil {
		panic("打开文件出错")
	}
	defer f.Close()
	return IsFileType(ft, f)
}

// IsFileType 判断指定文件与文件类型是否匹配
func IsFileType(ft *FileType, f *os.File) bool {
	size := len(ft.MagicNumber)
	for i := 0; i < size; i++ {
		bufSize := len(ft.MagicNumber[i])
		buf := make([]byte, bufSize)
		n, err := f.ReadAt(buf, 0)
		if err != nil {
			panic("判断文件类型读取魔数读取错误")
		}
		if n != bufSize {
			panic("判断文件类型读取魔数读取到的字节数目与期望不符合")
		}
		if IsByteArrayEqual(buf, ft.MagicNumber[i]) {
			return true
		}
	}
	return false
}

// CreateFileType 创建FileType
func CreateFileType(code string, name string, suffix []string, magicNumberStrings []string) *FileType {
	magicNumbers := make([][]byte, len(magicNumberStrings))
	for i := 0; i < len(magicNumberStrings); i++ {
		magicNumbers[i] = []byte(magicNumberStrings[i])
	}
	return &FileType{
		Code:        code,
		Name:        name,
		Suffix:      suffix,
		MagicNumber: magicNumbers,
	}
}

// 可执行文件

var WindowsExecuteFile = CreateFileType("WindowsExecuteFile", "Windows可执行文件", []string{"exe"}, []string{"MZ"})
var LinuxExecuteFile = CreateFileType("LinuxExecuteFile", "Linux可执行文件", []string{""}, []string{"\x7fELF"})
var JavaClassFile = CreateFileType("JavaClassFile", "Java类", []string{"class"}, []string{"\xCA\xFE\xBA\xBE"})

// 文档

var Office03File = CreateFileType("Office03File", "Office97-2003文件", []string{"doc"}, []string{"\xD0\xCF\x11\xE0\xA1\xB1\x1A\xE1"})
var PdfFile = CreateFileType("PdfFile", "pdf文件", []string{"pdf"}, []string{"%PDF-"})
var RtfFile = CreateFileType("RtfFile", "rtf文件", []string{"rtf"}, []string{"{\\rtf"})

// 压缩文件格式

var RarFile = CreateFileType("RarFile", "rar格式压缩文件", []string{"rar"}, []string{"Rar!"})
var X7zFile = CreateFileType("X7zFile", "7z格式压缩文件", []string{"7z"}, []string{"7z\xBC\xAF\x27"})
var ZipFile = CreateFileType("ZipFile", "zip格式压缩文件", []string{"zip"}, []string{"PK\x03\x04"})
var GzipFile = CreateFileType("GzipFile", "gzip格式压缩文件", []string{"gzip", "gz", "tgz"}, []string{"\x1F\x8B\x08"})
var RpmFile = CreateFileType("RpmFile", "rpm格式文件", []string{"rpm"}, []string{"\xED\xAB\xEE\xDB"})

// 图片文件格式

var JpegFile = CreateFileType("JpegFile", "jpeg图片格式", []string{"jpg", "jpeg"}, []string{"\xff\xd8\xff"})
var PngFile = CreateFileType("PngFile", "png图片格式", []string{"png"}, []string{"\x89PNG"})
var GifFile = CreateFileType("GifFile", "gif图片格式", []string{"gif"}, []string{"GIF"})
var BmpFile = CreateFileType("BmpFile", "bmp图片格式", []string{"bmp"}, []string{"BM"})
