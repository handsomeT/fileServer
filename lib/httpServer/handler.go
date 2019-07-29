package httpServer

import (
	"fileServer/cfg"
	"fileServer/lib/writer"
	"net/http"
)

func GetHandler(dir string) (handler func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// GET认为是文件目录
		if r.Method == "GET" {
			fileHandler(dir, path, w, r.Host)
			return
		}
		// Post认为是上传文件
		if r.Method == "POST" {
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
	}
}
