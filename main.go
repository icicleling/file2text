package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := os.Args[1]
	byte, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("err:", err)
	}
	str := base64.StdEncoding.EncodeToString(byte)
	fmt.Print(str)
}
