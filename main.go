package main

import (
	"encoding/json"
	"fileServer/cfg"
	"fileServer/lib/httpServer"
	"fileServer/lib/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// 读取配置文件
	d, err := ioutil.ReadFile(cfg.File)
	if err != nil {
		log.ErrorF("未获取到配置文件, 请确认文件存在：%s", cfg.File)
		return
	}
	var cfgData cfg.Cfg
	if err = json.Unmarshal(d, &cfgData); err != nil {
		log.Error("配置文件格式错误，请检查")
		return
	}
	http.HandleFunc("/", httpServer.GetHandler(cfgData.Dir))
	err = http.ListenAndServe(":"+strconv.Itoa(cfgData.Port), nil)
	if err != nil {
		log.Error("启动端口侦听失败")
		log.Error(err)
	}
}
