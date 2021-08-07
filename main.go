package main

import (
	"file2text/constants"
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println(time.Since(startTime))
	}()

	flag.Usage = func() {
		fmt.Print(constants.USAGE)
	}

	versionFlag := flag.BoolP("version", "v", false, "版本号")

	dataurlFlag := flag.Bool("dataurl", false, "转换为data url")
	printFlag := flag.Bool("print", false, "打印到终端")
	binFlag := flag.Bool("bin", false, "转换为二进制字符串")

	restoreFlag := flag.BoolP("restore", "r", false, "反转, 把base64字符串输出为文件")
	textFlag := flag.String("text", "", "base64字符串")

	flag.Parse()

	// 没有任何参数
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	// restore flag
	if *restoreFlag {
		Restore(textFlag, binFlag)
		return
	}

	// version flag
	if *versionFlag {
		fmt.Printf("v%s\n", constants.VERSION)
		return
	}

	// no flag
	Convert(dataurlFlag, printFlag, binFlag)
}
