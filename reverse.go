package main

import (
	"encoding/base64"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

func Reverse(textFlag *string, binFlag *bool) {
	originStr := ""
	filePath := "./output"
	resultByte := make([]byte, 0)

	if *textFlag != "" {
		originStr = *textFlag
		if flag.Arg(0) != "" {
			filePath = flag.Arg(0)
		}
	} else {
		pathStr := flag.Arg(0)
		byteArr, err := os.ReadFile(pathStr)
		if err != nil {
			log.Fatal(err)
		}
		originStr = string(byteArr)

		if flag.Arg(1) != "" {
			filePath = flag.Arg(1)
		} else {
			reg := regexp.MustCompile(`([^/\\\n]+)(?:\.[^/\\\n]+$)|([^/\\][^/\\\n.]+$)`)
			matchArr := reg.FindStringSubmatch(pathStr)
			hasExt := matchArr[1]
			noExt := matchArr[2]
			if hasExt != "" {
				filePath = "./" + hasExt
			} else if noExt != "" {
				filePath = "./" + noExt
			}
		}
	}

	if *binFlag {
		arr := strings.Split(originStr, " ")

		for i := 0; i < len(arr); i++ {
			num, err := strconv.ParseUint(arr[i], 2, 8)
			if err != nil {
				log.Fatal(err)
			}
			resultByte = append(resultByte, byte(num))
		}
	} else {
		result, err := base64.StdEncoding.DecodeString(originStr)
		if err != nil {
			log.Fatal(err)
		}
		resultByte = result
	}

	os.WriteFile(filePath, resultByte, 0666)
}
