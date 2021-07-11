package main

import (
	"fmt"
	"img2base64/util"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
	"regexp"

	flag "github.com/spf13/pflag"
)

func OutputFlag(dataUrlFlag *bool) {
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

	if *dataUrlFlag == true {
		ext := path.Ext(originPathStr)
		mimeType := mime.TypeByExtension(ext)
		base64Str = fmt.Sprintf("data:%s;base64,%s\n", mimeType, base64Str)
	}
	ioutil.WriteFile(targetPathStr, []byte(base64Str), 0666)
	os.Exit(0)
}
