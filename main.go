package main

import (
	"fmt"
	"img2base64/constants"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.Usage = func() {
		fmt.Print(constants.USAGE)
	}

	versionFlag := flag.BoolP("version", "v", false, "版本号")

	dataurlFlag := flag.Bool("dataurl", false, "转换为data url")
	printFlag := flag.Bool("print", false, "打印到终端")
	binFlag := flag.Bool("bin", false, "转换为二进制字符串")

	reverseFlag := flag.BoolP("reverse", "r", false, "反转, 把base64字符串输出为文件")
	textFlag := flag.String("text", "", "base64字符串")

	flag.Parse()

	// 没有任何参数
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	// reverse flag
	if *reverseFlag {
		Reverse(textFlag, binFlag)
		os.Exit(0)
	}

	// version flag
	if *versionFlag {
		fmt.Printf("v%s\n", constants.VERSION)
		os.Exit(0)
	}

	// no flag
	Output(dataurlFlag, printFlag, binFlag)
	os.Exit(0)
}
