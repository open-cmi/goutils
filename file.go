package goutils

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FileExist(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		return false
	}
	return fi != nil && !fi.IsDir()
}
