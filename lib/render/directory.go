package render

import (
	"os"
)

// 渲染目录
func Directory(baseUrl string, infoList []os.FileInfo) []byte {
	str := ""
	for _, info := range infoList {
		if isHiddenFile(info) {
			continue
		}
		str += linkStr(baseUrl, info.Name())
		str += "</br>"
	}
	return []byte(str)
}

// 是否是隐藏文件
func isHiddenFile(info os.FileInfo) bool {
	name := info.Name()
	if len(name) == 0 {
		return false
	}
	return name[0] == '.'
}