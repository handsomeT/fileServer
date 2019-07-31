package httpServer

import (
	"fileServer/lib/log"
	"fileServer/lib/util"
	"fileServer/lib/writer"
	"net/http"
)

func timeGetHandler(w http.ResponseWriter) {
	data, err := util.Cmd("date")
	if err != nil {
		log.Error(err)
		writer.InternalServerError(w)
		return
	}
	w.Write(data)
}

func timeSetHandler(time string, w http.ResponseWriter) {
	data, err := util.Cmd("sudo date -s \"" + time + "\"")
	if err != nil {
		log.Error(err)
		writer.InternalServerError(w)
		return
	}
	w.Write(data)
}

func timeRecHandler(w http.ResponseWriter) {
	data, err := util.Cmd("sudo ntpdate cn.pool.ntp.org")
	if err != nil {
		log.Error(err)
		writer.InternalServerError(w)
		return
	}
	w.Write(data)
}
