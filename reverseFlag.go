package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"regexp"
)

func ReverseFlag(pathFlag *string, textFlag *string) {
	base64Str := ""
	filePath := "./output"

	log.Println("pathflag: ", *pathFlag)

	if *pathFlag != "" {
		byte, err := ioutil.ReadFile(*pathFlag)
		if err != nil {
			log.Fatal(err)
		}
		base64Str = string(byte)

		if flag.Arg(0) != "" {
			filePath = flag.Arg(0)
		} else {
			reg := regexp.MustCompile(`([^/\\\n]+)(?:\.[^/\\\n]+$)|([^/\\][^/\\\n.]+$)`)
			matchArr := reg.FindStringSubmatch(*pathFlag)
			hasExt := matchArr[1]
			noExt := matchArr[2]
			if hasExt != "" {
				filePath = "./" + hasExt
			} else if noExt != "" {
				filePath = "./" + noExt
			}
		}
	} else if *textFlag != "" {
		base64Str = *textFlag
		if flag.Arg(0) != "" {
			filePath = flag.Arg(0)
		}
	} else {
		base64Str = flag.Arg(0)
		if flag.Arg(1) != "" {
			filePath = flag.Arg(1)
		}
	}

	result, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(filePath, result, 0666)
}
