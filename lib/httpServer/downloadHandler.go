package httpServer

import (
	"fileServer/lib/writer"
	"io"
	"net/http"
	"net/url"
	"os"
)

// 文件下载
func downloadHandler(dir string, path string, w http.ResponseWriter) {
	// 判断文件是否存在
	f, err := os.Open(dir + path)
	if err != nil {
		writer.NotFound(w)
		return
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil || len(fInfo.Name()) == 0 {
		writer.NotFound(w)
		return
	}
	// 处理目录
	if fInfo.IsDir() {
		writer.BadRequest(w)
		return
	}
	// 下载文件
	w.Header().Add("Content-Disposition", "attachment; filename="+url.PathEscape(fInfo.Name()))
	w.Header().Add("Content-Description", "File Transfer")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Expires", "0")
	w.Header().Add("Cache-Control", "must-revalidate")
	w.Header().Add("Pragma", "public")
	_, err = io.Copy(w, f)
	if err != nil {
		writer.InternalServerError(w)
	}
}
