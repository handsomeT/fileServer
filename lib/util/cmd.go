package util

import (
	"bytes"
	"os/exec"
)

func Cmd(data string) ([]byte, error) {
	cmd := exec.Command("/bin/bash", "-c", data)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

func Time() string {
	data, err := Cmd("date")
	if err != nil {
		return ""
	}
	return string(data)
}
