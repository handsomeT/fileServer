package httpServer

import (
	"fileServer/lib/writer"
	"io/ioutil"
	"net/http"
	"os"
)

// 文件处理
func fileHandler(dir string, path string, w http.ResponseWriter, host string) {
	// 读取文件
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
		list, err := f.Readdir(-1)
		if err != nil {
			writer.NotFound(w)
			return
		}
		// 这里需要传入路径来生成连接
		writer.Directory(w, host, path, list)
		return
	}
	// 处理文件
	data, err := ioutil.ReadFile(dir + path)
	if err != nil {
		writer.NotFound(w)
		return
	}
	writer.File(w, fInfo.Name(), data)
}
