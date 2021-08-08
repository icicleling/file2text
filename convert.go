package main

import (
	"encoding/base64"
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
	if *printFlag {
		defer fmt.Println()
	}

	originPathStr := flag.Arg(0)
	targetPathStr := flag.Arg(1)

	if targetPathStr == "" {
		reg := regexp.MustCompile(`[^/\\\n]+$`)
		matchArr := reg.FindStringSubmatch(originPathStr)
		fileName := matchArr[0]
		targetPathStr = "./" + fileName + ".txt"
	}

	var targetFile *os.File
	if !*printFlag {
		targetFileTemp, targetFileErr := os.Create(targetPathStr)
		if targetFileErr != nil {
			log.Fatal(targetFileErr)
		}
		defer targetFileTemp.Close()
		targetFile = targetFileTemp
	}

	originFile, originFileErr := os.Open(originPathStr)
	if originFileErr != nil {
		log.Fatal(originFileErr)
	}
	defer originFile.Close()

	if *binFlag {
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
			if *printFlag {
				fmt.Print(stringBuilder.String())
				continue
			}
			targetFile.WriteString(stringBuilder.String())
		}
		return
	}

	if *dataUrlFlag {
		ext := path.Ext(originPathStr)
		mimeType := mime.TypeByExtension(ext)
		if *printFlag {
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

		if *printFlag {
			fmt.Print(base64.StdEncoding.EncodeToString(p[:n]))
			continue
		}

		targetFile.WriteString(base64.StdEncoding.EncodeToString(p[:n]))
	}
}
