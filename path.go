package goutils

import (
	"os"
	"path"
	"strings"
)

var rootPath string = ""

// SetRootPath set root path, working directory
func SetRootPath(p string) {
	rootPath = p
	return
}

// GetRootPath get root path, working directory
func GetRootPath() string {
	if rootPath == "" {
		execFile, err := os.Executable()
		if err != nil {
			return ""
		}
		execPath := path.Dir(execFile)
		tmpdir := os.TempDir()
		if strings.HasPrefix(execFile, tmpdir) {
			execPath, err = os.Getwd()
			if err != nil {
				return ""
			}
		}

		rootPath = path.Dir(execPath)
	}
	return rootPath
}

// Getwd get pwd
func Getwd() string {
	execFile, err := os.Executable()
	if err != nil {
		return ""
	}
	execPath := path.Dir(execFile)
	tmpdir := os.TempDir()
	if strings.HasPrefix(execFile, tmpdir) {
		execPath, err = os.Getwd()
		if err != nil {
			return ""
		}
	}

	return execPath
}
