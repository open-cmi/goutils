package logutils

import (
	"log"

	"github.com/longbozhan/timewriter"
)

// 供全局调用的不同日志级别
var (
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

// Init init log
func Init(logPath string) {
	// 系统日志
	logWriter := GetLogWriter(logPath)
	Info = log.New(logWriter, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Error = log.New(logWriter, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Debug = log.New(logWriter, "[DEBUG] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

// GetLogWriter 获取系统日志配置
func GetLogWriter(logPath string) *timewriter.TimeWriter {
	wr := &timewriter.TimeWriter{
		Dir:        logPath, // 日志路径
		Compress:   true,    // 是否压缩
		ReserveDay: 30,      // 老化天数
	}

	return wr
}
