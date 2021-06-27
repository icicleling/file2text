package main

import (
	"flag"
	"img2base64/util"
	"io/ioutil"
	"log"
	"regexp"
)

func OutputFlag() {
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
	ioutil.WriteFile(targetPathStr, []byte(base64Str), 0666)
}
