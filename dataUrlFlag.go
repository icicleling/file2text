package main

import (
	"flag"
	"fmt"
	"img2base64/util"
	"log"
	"mime"
	"os"
	"path"
)

func DataUrlFlag() {
	pathStr := flag.Arg(0)
	ext := path.Ext(pathStr)
	base64Str, err := util.GetBase64(pathStr)
	if err != nil {
		log.Fatal(err)
	}

	mimeType := mime.TypeByExtension(ext)
	dataurl := fmt.Sprintf("data:%s;base64,%s\n", mimeType, base64Str)

	fmt.Printf(dataurl)
	os.Exit(0)
}
