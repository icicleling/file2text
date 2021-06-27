package util

import (
	"encoding/base64"
	"io/ioutil"
)

// 读取路径文件转换后, 返回base64字符串
func GetBase64ByFilePath(pathStr string) (string, error) {
	byte, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(byte)
	return base64Str, nil
}
