package writer

import (
	"fileServer/lib/log"
	"net/http"
)

func write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
	}
}
