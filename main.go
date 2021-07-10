package main

import (
	"flag"
	"fmt"
	"img2base64/constants"
	"img2base64/util"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Print(constants.USAGE)
	}

	versionFlag := flag.Bool("version", false, "版本号")
	outputFlag := flag.Bool("output", false, "输出文件")
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
		base64Str, err := util.GetBase64ByFilePath(pathStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(base64Str)
	}

	// output flag
	if *outputFlag {
		OutputFlag(dataurlFlag)
	}

	// reverse flag
	if *reverseFlag {
		ReverseFlag(pathFlag, textFlag)
	}

	// version flag
	if *versionFlag {
		fmt.Printf("v%s\n", constants.VERSION)
		os.Exit(0)
	}
}
