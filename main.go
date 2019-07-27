package main

import (
	"encoding/json"
	"fileServer/cfg"
	"fileServer/lib/log"
	"fileServer/lib/render"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 目录
var dir string

func getFileFormat(fileName string) string {
	if len(fileName) == 0 {
		return ""
	}
	i := strings.LastIndex(fileName, ".")
	if i < 0 {
		return ""
	}
	return fileName[i:]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// 完整目录
	fullPath := dir + path
	// 读取文件
	f, err := os.Open(fullPath)
	if err != nil {
		w.Write(render.NotFount())
		return
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil || len(fInfo.Name()) == 0 {
		w.Write(render.NotFount())
		return
	}
	if fInfo.IsDir() {
		list, err := f.Readdir(-1)
		if err != nil {
			w.Write(render.NotFount())
			return
		}
		// 这里需要传入路径来生成连接
		baseUrl := r.Host + path
		w.Write(render.Directory(baseUrl, list))
		return
	}
	// 读取文件数据
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		w.Write(render.NotFount())
		return
	}
	//
	switch getFileFormat(fInfo.Name()) {
	case ".md":
		w.Write(render.Markdown(data))
	default:
		w.Write(data)
	}
}

func main() {
	// 读取配置文件
	d, err := ioutil.ReadFile(cfg.File)
	if err != nil {
		log.ErrorF("未获取到配置文件, 请确认文件存在：%s", cfg.File)
		return
	}
	var cfgData cfg.Cfg
	if err = json.Unmarshal(d, &cfgData); err != nil {
		log.Error("配置文件格式错误，请检查")
		return
	}
	dir = cfgData.Dir
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":" + strconv.Itoa(cfgData.Port), nil)
}