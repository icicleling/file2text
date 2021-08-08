package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

func Restore(textFlag *string, binFlag *bool) {
	originPathStr := flag.Arg(0)
	targetPathStr := "./output"

	if *textFlag != "" {
		if flag.Arg(0) != "" {
			targetPathStr = flag.Arg(0)
		}
	} else {
		if flag.Arg(1) != "" {
			targetPathStr = flag.Arg(1)
		} else {
			reg := regexp.MustCompile(`([^/\\\n]+)(?:\.[^/\\\n]+$)|([^/\\][^/\\\n.]+$)`)
			matchArr := reg.FindStringSubmatch(originPathStr)
			hasExt := matchArr[1]
			noExt := matchArr[2]
			if hasExt != "" {
				targetPathStr = "./" + hasExt
			} else if noExt != "" {
				targetPathStr = "./" + noExt
			}
		}
	}

	originFile, originFileErr := os.Open(originPathStr)
	if originFileErr != nil {
		log.Fatal(originFileErr)
	}
	defer originFile.Close()

	targetFile, targetFileErr := os.Create(targetPathStr)
	if targetFileErr != nil {
		log.Fatal(targetFileErr)
	}
	defer targetFile.Close()

	if *binFlag {
		var reader io.Reader = originFile

		if *textFlag != "" {
			reader = strings.NewReader(*textFlag)
		}

		p := make([]byte, 8192)
		for {
			n, err := reader.Read(p)
			if err != nil || err == io.EOF {
				break
			}

			pLen := len(p[:n])
			resultBytes := make([]byte, pLen/8)
			for i := 0; i < pLen/8; i++ {
				num, err := strconv.ParseUint(string(p[i*8:i*8+8]), 2, 8)
				if err != nil {
					log.Fatal(err)
				}
				resultBytes[i] = byte(num)
			}
			targetFile.Write(resultBytes)
		}
		return
	}

	if *textFlag != "" {
		resultBytes, err := base64.StdEncoding.DecodeString(*textFlag)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(targetPathStr, resultBytes, 0666)
		return
	}

	p := make([]byte, 4*1024)
	for {
		n, nErr := originFile.Read(p)
		if nErr != nil || nErr == io.EOF {
			break
		}

		byteArr, byteArrErr := base64.StdEncoding.DecodeString(string(p[:n]))
		if byteArrErr != nil {
			log.Fatal(byteArrErr)
		}

		targetFile.Write(byteArr)
	}
}
