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
	"regexp"
)

// VERSION 版本号
const VERSION = "0.4.1"

// USAGE help说明文本
const USAGE = `
转换并打印base64字符串
	img2base64 <path>

flag:
	-dataurl <path>
		输出dataurl

	-reverse [-path|-text] <base64-string|base64filepath> [filepath]
		反转base64字符串为文件, 输入字符串或路径, 默认为字符串
		注意命令行是有长度限制的, 不同终端限制长度不同
		不写filepath的话, -text默认输出名为'output'的文件, -path默认输出去掉后缀的同名文件

	-version
		打印版本号
`

func main() {
	flag.Usage = func() {
		fmt.Print(USAGE)
	}

	versionFlag := flag.Bool("version", false, "版本号")
	dataurlFlag := flag.Bool("dataurl", false, "输出data url")
	reverseFlag := flag.Bool("reverse", false, "反转, 把base64字符串输出为文件")
	pathFlag := flag.String("path", "", "文本文件路径")
	textFlag := flag.String("text", "", "base64字符串")
	flag.Parse()

	// 没有任何参数
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
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
		base64Str := ""
		filePath := "./output"
		if *pathFlag != "" {
			byte, err := ioutil.ReadFile(*pathFlag)
			if err != nil {
				log.Fatal(err)
			}
			base64Str = string(byte)

			if flag.Arg(0) != "" {
				filePath = flag.Arg(0)
			} else {
				reg := regexp.MustCompile(`([^/\\\n]+)(?:\.[^/\\\n]+$)|([^/\\][^/\\\n.]+$)`)
				matchArr := reg.FindStringSubmatch(*pathFlag)
				hasExt := matchArr[1]
				noExt := matchArr[2]
				if hasExt != "" {
					filePath = "./" + hasExt
				} else if noExt != "" {
					filePath = "./" + noExt
				}
			}
		} else if *textFlag != "" {
			base64Str = *textFlag
			if flag.Arg(0) != "" {
				filePath = flag.Arg(0)
			}
		} else {
			base64Str = flag.Arg(0)
			if flag.Arg(1) != "" {
				filePath = flag.Arg(1)
			}
		}

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
