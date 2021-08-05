# 更新日志

## 0.11
- feat: 
  - go mod 更新go版本到1.16, ioutil换成os
  - reverse命令改名为restore

## 0.10.1

- feat: 使用 strings.builder 拼接字符串, 极大幅优化了转换二进制字符串的性能

## 0.10

- feat: 改名叫 file2text

## 0.9

- feat:
  - 添加--bin 命令, 输出和反转二进制格式字符串
  - 简化命令 --reverse 移除默认的 path 命令

## 0.8

- feat:
  - 简化命令, 增加缩写命令
  - 移除 output 命令, 默认行为使用原 output 命令的效果

## 0.7

- feat:
  - 修改--release 默认行为
  - dataurl 选项移入 output 子选项

## 0.6

- feat: 添加 --output 输出结果到文本文件

## 0.5.1

- refactor: 拆分一下文件

## 0.5

- feat: --reverse 允许不写目标路径, 可以生成默认文件名

## 0.4.1

- fix: 修复没有任何参数时显示的 help 信息不正确的问题

## 0.4

- feat:
  - 更好的 help 显示文档
  - 参数-reverse, 把 base64 字符串反转回文件

## 0.3.1

- docs: 忘了写更新日志, 写上

## 0.3

- feat:
  - 增加参数 -version -dataurl
    - -version 打印版本号
    - -dataurl 输出 dataurl 格式的 base64 字符串
  - 默认打印纯 base64 字符串

## 0.2

- feat: 默认打印 base64 data url, 可以直接放 markdown 里用

## 0.1

- feat: 输入图片路径, 打印 base64 字符串
