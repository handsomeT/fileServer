package writer

import (
	"fileServer/lib/log"
	"html/template"
	"net/http"
	"os"
)

type Item struct {
	Href   string
	Name   string
	IsFile bool
}

// 渲染目录
func Directory(w http.ResponseWriter, host string, dir string, infoList []os.FileInfo) {
	var list []Item
	for _, info := range infoList {
		// 过滤隐藏文件
		if isHiddenFile(info) {
			continue
		}
		item := Item{
			Href:   host + dir + "/" + info.Name(),
			Name:   info.Name(),
			IsFile: !info.IsDir(),
		}
		list = append(list, item)
	}
	// 从模板中生成
	t, _ := template.ParseFiles("template/directory.html")
	data := make(map[string]interface{})
	data["list"] = list
	// 添加上传url
	data["uploadUrl"] = host + dir
	err := t.Execute(w, data)
	if err != nil {
		log.Error(err)
		InternalServerError(w)
	}
}

// 是否是隐藏文件
func isHiddenFile(info os.FileInfo) bool {
	name := info.Name()
	if len(name) == 0 {
		return false
	}
	return name[0] == '.'
}
