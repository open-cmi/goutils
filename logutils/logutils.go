package logutils

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
)

// Logger log logger
var Logger *log.Logger

var LogFullPath string = ""

func SetLogOption(p string) {
	LogFullPath = p
	return
}

func GetLogger() *log.Logger {
	if Logger == nil {
		if LogFullPath == "" {
			executable, _ := os.Executable()
			procname := path.Base(executable)
			rp := common.GetRootPath()
			LogFullPath = filepath.Join(rp, "data", procname+".log")
			logDir := filepath.Join(rp, "data")
			if !goutils.IsExist(logDir) {
				os.MkdirAll(logDir, os.ModePerm)
			}
		}

		w, err := os.OpenFile(LogFullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return nil
		}
		logger := log.New(w, "", log.LstdFlags)
		Logger = logger
	}

	return Logger
}
