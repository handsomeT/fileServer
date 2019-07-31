package util

import (
	"bytes"
	"fileServer/lib/log"
	"os/exec"
)

func Cmd(data string) ([]byte, error) {
	log.Debug(data)
	cmd := exec.Command("/bin/bash", "-c", data)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}
