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
		// 如果请求方法不是get
		if r.Method != "GET" {
			writer.MethodNotAllowed(w)
			return
		}
		// redirect / 到 文件根目录
		path := r.URL.Path
		if len(path) == 0 || path == "/" {
			http.Redirect(w, r, cfg.UrlFile, http.StatusFound)
			return
		}
		// 文件目录
		if flag, path := removeIndex(path, cfg.UrlFile); flag {
			fileHandler(dir, path, w, r.Host)
			return
		}
		// 下载
		if flag, path := removeIndex(path, cfg.UrlDownload); flag {
			downloadHandler(dir, path, w)
			return
		}
		// 上传处理
		if flag, _ := removeIndex(path, cfg.UrlUpload); flag {
			uploadHandler()
			return
		}
	}
}
