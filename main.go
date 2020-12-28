package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
)

// VERSION 版本号
const VERSION = "0.3"

func main() {
	versionFlag := flag.Bool("version", false, "版本号")
	dataurlFlag := flag.Bool("dataurl", false, "输出data url")
	reverseFlag := flag.Bool("reverse", false, "反转, 把base64字符串输出为文件")
	flag.Parse()

	// 没有任何参数
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 没有flag 有其他参数
	if flag.NFlag() == 0 {
		pathStr := flag.Arg(0)
		base64Str, err := getBase64(pathStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(base64Str)
	}

	// dataurl flag
	if *dataurlFlag {
		pathStr := flag.Arg(0)
		ext := path.Ext(pathStr)
		base64Str, err := getBase64(pathStr)
		if err != nil {
			log.Fatal(err)
		}

		mimeType := mime.TypeByExtension(ext)
		dataurl := fmt.Sprintf("data:%s;base64,%s\n", mimeType, base64Str)

		fmt.Printf(dataurl)
		os.Exit(0)
	}

	// reverse flag
	if *reverseFlag {
		base64Str := flag.Arg(0)
		filePath := flag.Arg(1)
		result, err := base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(filePath, result, 0666)
	}

	// version flag
	if *versionFlag {
		fmt.Printf("v%s\n", VERSION)
		os.Exit(1)
	}
}

// 读取路径文件转换后, 返回base64字符串
func getBase64(pathStr string) (string, error) {
	byte, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(byte)
	return base64Str, nil
}
