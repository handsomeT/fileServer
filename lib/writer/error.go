package writer

import (
	"net/http"
)

// 404 not found
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	write(w, []byte("404 not found"))
}

// 不允许的请求方法
func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	write(w, []byte("不受支持的请求方法"))
}

// 服务器内部错误
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	write(w, []byte("抱歉, 服务器出现了一个错误"))
}

// bad request
func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	write(w, []byte("bad request"))
}
