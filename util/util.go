package util

import (
	"encoding/base64"
	"os"
)

// 读取路径文件, 返回base64字符串
func GetBase64ByFilePath(pathStr string) (string, error) {
	byte, err := os.ReadFile(pathStr)
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(byte)
	return base64Str, nil
}

// 读取路径文件, 返回byte
func GetByteByFilePath(pathStr string) ([]byte, error) {
	byte, err := os.ReadFile(pathStr)
	if err != nil {
		return nil, err
	}
	return byte, nil
}
