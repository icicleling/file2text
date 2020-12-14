package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
)

func main() {
	pathStr := os.Args[1]
	ext := path.Ext(pathStr)
	byte, err := ioutil.ReadFile(pathStr)
	if err != nil {
		log.Fatal(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(byte)
	mimeType := mime.TypeByExtension(ext)

	fmt.Printf("data:%s;base64,%s", mimeType, base64Str)
}
