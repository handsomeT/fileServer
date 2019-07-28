package httpServer

import (
	"fileServer/lib/writer"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// 上传处理
func uploadHandler(dir string, path string, f multipart.File, fHeader *multipart.FileHeader, w http.ResponseWriter) {
	// 读取文件数据
	data, err := ioutil.ReadAll(f)
	if err != nil {
		writer.BadRequest(w)
		return
	}
	// 保存文件
	savePath := dir + path + "/" + fHeader.Filename
	saveFile, err := os.Create(savePath)
	if err != nil {
		writer.InternalServerError(w)
		return
	}
	defer saveFile.Close()
	if _, err := saveFile.Write(data); err != nil {
		writer.InternalServerError(w)
		return
	}
	_, err = w.Write([]byte("上传成功"))
	if err != nil {
		writer.InternalServerError(w)
		return
	}
}
