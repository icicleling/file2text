package constants

// USAGE help说明文本
const USAGE = `
转换并打印base64字符串
	img2base64 <path>

flag:
	-output <filepath> [target-filepath]
		转换后的文本输出一个文本文件, 默认为filename.ext.txt

	-dataurl <filepath>
		输出dataurl

	-reverse [-path|-text] <base64-string|base64filepath> [target-filepath]
		反转base64字符串为文件, 输入路径(path)或字符串(text), 默认为path
		注意命令行是有长度限制的, 不同终端限制长度不同
		不写filepath的话, -text默认输出名为'output'的文件, -path默认输出去掉后缀的同名文件

	-version
		打印版本号
`

// VERSION 版本号
const VERSION = "0.6"
