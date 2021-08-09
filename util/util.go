package util

import "regexp"

func GetPathAppendExt(originPath string) string {
	reg := regexp.MustCompile(`[^/\\\n]+$`)
	matchArr := reg.FindStringSubmatch(originPath)
	fileName := matchArr[0]
	return "./" + fileName + ".txt"
}

func GetPathRemoveExt(originPath string) string {
	reg := regexp.MustCompile(`([^/\\\n]+)(?:\.[^/\\\n]+$)|([^/\\][^/\\\n.]+$)`)
	matchArr := reg.FindStringSubmatch(originPath)
	hasExt := matchArr[1]
	noExt := matchArr[2]

	if hasExt != "" {
		return "./" + hasExt
	}
	return "./" + noExt
}
