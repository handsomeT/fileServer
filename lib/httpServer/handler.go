package httpServer

import (
	"fileServer/lib/writer"
	"fileserver/cfg"
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
		// 如果是获取时间
		if flag, _ := removePrefix(path, cfg.UrlTimeGet); flag {
			// 获取时间
			data, err := getTime()
			if err != nil {
				writer.InternalServerError(w)
				return
			}
			w.Write(data)
		}
	}
}
