package main

import (
	"fmt"
	"img2base64/util"
	"io/ioutil"
	"log"
	"mime"
	"path"
	"regexp"

	flag "github.com/spf13/pflag"
)

func Output(dataUrlFlag *bool, printFlag *bool) {
	originPathStr := flag.Arg(0)
	targetPathStr := flag.Arg(1)

	if targetPathStr == "" {
		reg := regexp.MustCompile(`[^/\\\n]+$`)
		matchArr := reg.FindStringSubmatch(*&originPathStr)
		fileName := matchArr[0]
		targetPathStr = "./" + fileName + ".txt"
	}

	base64Str, err := util.GetBase64ByFilePath(originPathStr)
	if err != nil {
		log.Fatal(err)
	}

	if *dataUrlFlag {
		ext := path.Ext(originPathStr)
		mimeType := mime.TypeByExtension(ext)
		base64Str = fmt.Sprintf("data:%s;base64,%s\n", mimeType, base64Str)
	}
	if *printFlag {
		fmt.Println(base64Str)
		return
	}
	ioutil.WriteFile(targetPathStr, []byte(base64Str), 0666)
}
