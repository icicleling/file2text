package main

import (
	"file2text/util"
	"fmt"
	"log"
	"mime"
	"os"
	"path"
	"regexp"
	"strings"

	flag "github.com/spf13/pflag"
)

func Output(dataUrlFlag *bool, printFlag *bool, binFlag *bool) {
	originPathStr := flag.Arg(0)
	targetPathStr := flag.Arg(1)
	resultStr := ""

	if targetPathStr == "" {
		reg := regexp.MustCompile(`[^/\\\n]+$`)
		matchArr := reg.FindStringSubmatch(originPathStr)
		fileName := matchArr[0]
		targetPathStr = "./" + fileName + ".txt"
	}

	if *binFlag {
		byteArr, err := util.GetByteByFilePath(originPathStr)
		if err != nil {
			log.Fatal(err)
		}
		var stringBuilder strings.Builder
		for i := 0; i < len(byteArr); i++ {
			stringBuilder.WriteString(fmt.Sprintf("%b ", byteArr[i]))
		}
		resultStr = strings.TrimSpace(stringBuilder.String())
	} else {
		base64Str, err := util.GetBase64ByFilePath(originPathStr)
		if err != nil {
			log.Fatal(err)
		}
		resultStr = base64Str
	}

	if *dataUrlFlag && !*binFlag {
		ext := path.Ext(originPathStr)
		mimeType := mime.TypeByExtension(ext)
		resultStr = fmt.Sprintf("data:%s;base64,%s\n", mimeType, resultStr)
	}

	if *printFlag {
		fmt.Println(resultStr)
		return
	}
	os.WriteFile(targetPathStr, []byte(resultStr), 0666)
}
