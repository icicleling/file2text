package util

import (
	"os"
)

// 读取路径文件, 返回byte
func GetByteByFilePath(pathStr string) ([]byte, error) {
	byte, err := os.ReadFile(pathStr)
	if err != nil {
		return nil, err
	}
	return byte, nil
}
