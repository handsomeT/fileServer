package httpServer

import (
	"fileServer/cfg"
	"fileServer/lib/writer"
	"net/http"
	"strings"
)

func removePrefix(url string, prefix string) (bool, string) {
	if strings.Index(url, prefix) != 0 {
		return false, ""
	}
	return true, url[len(prefix):]
}

func GetHandler(dir string) (handler func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// GET认为是文件目录
		if r.Method == "GET" {
			fileHandler(dir, path, w, r.Host)
			return
		}
		// 暂不支持GET和POST以外的请求方式
		if r.Method != "POST" {
			writer.MethodNotAllowed(w)
			return
		}
		// 判断是否是上传文件
		if flag, path := removePrefix(path, cfg.UrlUpload); flag {
			// 判断文件大小
			if err := r.ParseMultipartForm(cfg.MaxUploadSize); err != nil {
				writer.BadRequest(w)
				return
			}
			// 读取文件数据
			f, fHeader, err := r.FormFile("file")
			if err != nil {
				writer.BadRequest(w)
				return
			}
			defer f.Close()
			// 上传文件
			uploadHandler(dir, path, f, fHeader, w)
			return
		}
		// 时间相关处理
		if flag, _ := removePrefix(path, cfg.UrlTimeGet); flag {
			timeGetHandler(w)
			return
		}
		if flag, _ := removePrefix(path, cfg.UrlTimeSet); flag {
			if err := r.ParseMultipartForm(cfg.MaxUploadSize); err != nil {
				writer.BadRequest(w)
				return
			}
			time := r.FormValue("time")
			if len(time) == 0 {
				writer.BadRequest(w)
				return
			}
			timeSetHandler("", w)
			return
		}
		if flag, _ := removePrefix(path, cfg.UrlTimeRec); flag {
			timeRecHandler(w)
			return
		}
	}
}
