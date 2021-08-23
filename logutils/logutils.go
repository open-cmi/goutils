package logutils

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
)

// Logger log logger
var Logger *log.Logger

var LogFullPath string = ""
var LogDir string = ""

func SetLogDir(p string) {
	LogDir = p
}

// FormatLogPath format path
func FormatLogPath(t *time.Time) string {
	executable, _ := os.Executable()
	procname := path.Base(executable)
	newfile := fmt.Sprintf("%s-%d-%d-%d.log", procname, t.Year(), t.Month(), t.Day())
	if LogDir == "" {
		rp := common.GetRootPath()
		LogDir = filepath.Join(rp, "log")
	}

	if !goutils.IsExist(LogDir) {
		os.MkdirAll(LogDir, os.ModePerm)
	}

	fullpath := filepath.Join(LogDir, newfile)
	return fullpath
}

// Ticker ticker
func Ticker() {
	// 每半小时检测一下时间
	ticker := time.NewTicker(30 * 60 * time.Second)
	defer ticker.Stop()
	preday := time.Now().Day()
	for cur := range ticker.C {
		if preday != cur.Day() {
			// create new file
			LogFullPath = FormatLogPath(&cur)
			w, err := os.OpenFile(LogFullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
			if err != nil {
				continue
			}
			Logger.SetOutput(w)
			preday = cur.Day()
		}
	}
}

func GetLogger() *log.Logger {
	if Logger == nil {
		if LogFullPath == "" {
			t := time.Now()
			LogFullPath = FormatLogPath(&t)
		}

		w, err := os.OpenFile(LogFullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return nil
		}
		logger := log.New(w, "", log.LstdFlags)
		Logger = logger
		go Ticker()
	}

	return Logger
}
