package sysutil

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetHostName() string {
	if runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
		rf, err := os.Open("/etc/hostname")
		if err != nil {
			return ""
		}
		bt, err := ioutil.ReadAll(rf)
		if err != nil {
			return ""
		}
		return strings.Trim(string(bt), "\n\t ")
	} else if runtime.GOOS == "darwin" {
		args := []string{"--get", "LocalHostName"}
		bt, err := exec.Command("scutil", args...).Output()
		if err != nil {
			return ""
		}
		return strings.Trim(string(bt), "\n\t ")
	}
	return ""
}
