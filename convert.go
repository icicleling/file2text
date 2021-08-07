package main

import (
	"encoding/base64"
	"file2text/util"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path"
	"regexp"
	"strings"

	flag "github.com/spf13/pflag"
)

func Convert(dataUrlFlag *bool, printFlag *bool, binFlag *bool) {
	originPathStr := flag.Arg(0)
	targetPathStr := flag.Arg(1)

	if targetPathStr == "" {
		reg := regexp.MustCompile(`[^/\\\n]+$`)
		matchArr := reg.FindStringSubmatch(originPathStr)
		fileName := matchArr[0]
		targetPathStr = "./" + fileName + ".txt"
	}

	targetFile, targetFileErr := os.Create(targetPathStr)
	if targetFileErr != nil {
		log.Fatal(targetFileErr)
	}
	defer targetFile.Close()

	originFile, originFileErr := os.Open(originPathStr)
	if originFileErr != nil {
		log.Fatal(originFileErr)
	}
	defer originFile.Close()

	if *binFlag {
		byteArr, err := util.GetByteByFilePath(originPathStr)
		if err != nil {
			log.Fatal(err)
		}
		var stringBuilder strings.Builder
		for i := 0; i < len(byteArr); i++ {
			stringBuilder.WriteString(fmt.Sprintf("%b ", byteArr[i]))
		}
		resultStr := strings.TrimSpace(stringBuilder.String())
		if *printFlag {
			fmt.Println(resultStr)
		}
		os.WriteFile(targetPathStr, []byte(resultStr), 0666)
		return
	}

	if *dataUrlFlag {
		fmt.Println("dataurl")
		ext := path.Ext(originPathStr)
		mimeType := mime.TypeByExtension(ext)
		targetFile.Write([]byte(fmt.Sprintf("data:%s;base64,", mimeType)))
	}

	p := make([]byte, 3*1024)
	for {
		n, err := originFile.Read(p)
		if err != nil || err == io.EOF {
			if *printFlag {
				fmt.Print("\n")
			}
			break
		}

		if *printFlag {
			fmt.Print(base64.StdEncoding.EncodeToString(p[:n]))
			continue
		}

		targetFile.Write([]byte(base64.StdEncoding.EncodeToString(p[:n])))
	}
}
