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
	"strings"

	flag "github.com/spf13/pflag"
)

func Convert(dataUrlFlag bool, printFlag bool, binFlag bool) {
	if printFlag {
		defer fmt.Println()
	}

	originFile := os.NewFile(0, "")
	targetFile := os.NewFile(0, "")

	setConvertFileVar(originFile, targetFile, printFlag)
	defer originFile.Close()
	defer targetFile.Close()

	if binFlag {
		convertBinary(originFile, targetFile, printFlag)
		return
	}

	convertBase64(originFile, targetFile, printFlag, dataUrlFlag)
}

func convertBinary(originFile *os.File, targetFile *os.File, printFlag bool) {
	p := make([]byte, 8192)
	for {
		n, err := originFile.Read(p)
		if err != nil || err == io.EOF {
			break
		}

		pLen := len(p[:n])
		var stringBuilder strings.Builder
		for i := 0; i < pLen; i++ {
			stringBuilder.WriteString(fmt.Sprintf("%0.8b", p[i]))
		}
		if printFlag {
			fmt.Print(stringBuilder.String())
			continue
		}
		targetFile.WriteString(stringBuilder.String())
	}
}

func convertBase64(originFile *os.File, targetFile *os.File, printFlag bool, dataUrlFlag bool) {
	if dataUrlFlag {
		fileinfo, err := originFile.Stat()
		if err != nil {
			log.Fatal(err)
		}
		ext := path.Ext(fileinfo.Name())
		mimeType := mime.TypeByExtension(ext)
		if printFlag {
			fmt.Printf("data:%s;base64,", mimeType)
		} else {
			targetFile.WriteString(fmt.Sprintf("data:%s;base64,", mimeType))
		}
	}

	p := make([]byte, 3*1024)
	for {
		n, err := originFile.Read(p)
		if err != nil || err == io.EOF {
			break
		}

		if printFlag {
			fmt.Print(base64.StdEncoding.EncodeToString(p[:n]))
			continue
		}

		targetFile.WriteString(base64.StdEncoding.EncodeToString(p[:n]))
	}
}

func setConvertFileVar(originFile *os.File, targetFile *os.File, printFlag bool) {
	originPathStr := flag.Arg(0)
	targetPathStr := flag.Arg(1)

	originFileT, originFileErr := os.Open(originPathStr)
	if originFileErr != nil {
		log.Fatal(originFileErr)
	}
	*originFile = *originFileT

	if targetPathStr == "" {
		targetPathStr = util.GetPathAppendExt(originPathStr)
	}

	if !printFlag {
		targetFileT, targetFileErr := os.Create(targetPathStr)
		if targetFileErr != nil {
			log.Fatal(targetFileErr)
		}
		*targetFile = *targetFileT
	}
}
