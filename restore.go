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

	if *binFlag {
		originStr := *textFlag
		if *textFlag == "" {
			byteArr, err := os.ReadFile(originPathStr)
			if err != nil {
				log.Fatal(err)
			}
			originStr = string(byteArr)
		}

		arr := strings.Split(originStr, " ")
		resultBytes := make([]byte, 0)

		for i := 0; i < len(arr); i++ {
			num, err := strconv.ParseUint(arr[i], 2, 8)
			if err != nil {
				log.Fatal(err)
			}
			resultBytes = append(resultBytes, byte(num))
		}
		os.WriteFile(targetPathStr, resultBytes, 0666)
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
