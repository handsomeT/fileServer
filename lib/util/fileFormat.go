package util

import "strings"

func GetFileFormat(fileName string) string {
	if len(fileName) == 0 {
		return ""
	}
	i := strings.LastIndex(fileName, ".")
	if i < 0 {
		return ""
	}
	return fileName[i:]
}
