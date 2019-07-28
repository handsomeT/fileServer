package httpServer

import (
	"fileServer/cfg"
	"fileServer/lib/writer"
	"net/http"
	"strings"
)

func removeIndex(url string, index string) (bool, string) {
	if strings.Index(url, index) != 0 {
		return false, ""
	}
	return true, url[len(index):]
}

func GetHandler(dir string) (handler func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		// redirect / 到 文件根目录
		path := r.URL.Path
		if len(path) == 0 || path == "/" {
			http.Redirect(w, r, cfg.UrlFile, http.StatusFound)
			return
		}
		// 文件目录
		if flag, path := removeIndex(path, cfg.UrlFile); flag {
			// 如果请求方法不是get
			if r.Method != "GET" {
				writer.MethodNotAllowed(w)
				return
			}
			fileHandler(dir, path, w, r.Host)
			return
		}
		// 下载
		if flag, path := removeIndex(path, cfg.UrlDownload); flag {
			// 如果请求方法不是get
			if r.Method != "GET" {
				writer.MethodNotAllowed(w)
				return
			}
			downloadHandler(dir, path, w)
			return
		}
		// 上传处理
		if flag, path := removeIndex(path, cfg.UrlUpload); flag {
			// 文件上传要求方法是Post
			if r.Method != "POST" {
				writer.MethodNotAllowed(w)
				return
			}
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
