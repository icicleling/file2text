package constants

// USAGE help说明文本
const USAGE = `
转换并打印base64字符串
	img2base64 <path>

flag:
	-output [--dataurl] <filepath> [target-filepath]
		转换文件为文本文件, 输出文件名格式为 filename.ext.txt
		--dataurl 输出dataurl

	-reverse [-path|-text] <base64-string|base64filepath> [target-filepath]
		反转base64字符串为文件
		--path 默认值, 从读取路径的文件, 输出去掉后缀的同名文件
		--text 读取命令行输入的文本, 输出名为'output'的文件
		注意命令行是有长度限制的, 不同终端限制长度不同

	-version
		打印版本号
`

// VERSION 版本号
const VERSION = "0.6"
