package main

import (
	"encoding/base64"
	"file2text/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

func Restore(textFlag string, binFlag bool) {
	originFile := os.NewFile(0, "")
	targetFile := os.NewFile(0, "")

	setRestoreFileVar(originFile, targetFile, textFlag)
	defer originFile.Close()
	defer targetFile.Close()

	if binFlag {
		restoreBinary(originFile, targetFile, textFlag)
		return
	}

	restoreBase64(originFile, targetFile, textFlag)
}

func restoreBinary(originFile *os.File, targetFile *os.File, textFlag string) {
	var reader io.Reader = originFile

	if textFlag != "" {
		reader = strings.NewReader(textFlag)
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
}

func restoreBase64(originFile *os.File, targetFile *os.File, textFlag string) {
	if textFlag != "" {
		resultBytes, err := base64.StdEncoding.DecodeString(textFlag)
		if err != nil {
			log.Fatal(err)
		}
		targetFile.Write(resultBytes)
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

func setRestoreFileVar(originFile *os.File, targetFile *os.File, textFlag string) {
	targetFilePath := "./output"
	if textFlag != "" {
		if flag.Arg(0) != "" {
			targetFilePath = flag.Arg(0)
		}
	} else {
		// origin file
		originPathStr := flag.Arg(0)
		originFileTemp, originFileErr := os.Open(originPathStr)
		if originFileErr != nil {
			log.Fatal(originFileErr)
		}
		*originFile = *originFileTemp

		// target file
		if flag.Arg(1) != "" {
			targetFilePath = flag.Arg(1)
		} else {
			targetFilePath = util.GetPathRemoveExt(originPathStr)
		}
	}
	targetFileTemp, err := os.Create(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	*targetFile = *targetFileTemp
}
