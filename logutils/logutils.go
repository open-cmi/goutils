package logutils

import (
	"io"
	"log"

	"github.com/longbozhan/timewriter"
)

const (
	Debug = 0
	Info  = 1
	Error = 2
)

type Logger struct {
	logger   map[int]*log.Logger
	LogLevel int
	Writer   io.Writer
}

// 供全局调用的不同日志级别
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// Init init log
func Init(logPath string) {
	// 系统日志
	logWriter := GetLogWriter(logPath)
	InfoLogger = log.New(logWriter, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	ErrorLogger = log.New(logWriter, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	DebugLogger = log.New(logWriter, "[DEBUG] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
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

func NewLogger(logPath string) *Logger {
	var logger Logger

	wr := timewriter.TimeWriter{
		Dir:        logPath, // 日志路径
		Compress:   true,    // 是否压缩
		ReserveDay: 30,      // 老化天数
	}

	logger.logger = make(map[int]*log.Logger, 0)
	logger.logger[Info] = log.New(&wr, "[INFO] ", log.Ldate|log.Lmicroseconds)
	logger.logger[Error] = log.New(&wr, "[ERROR] ", log.Ldate|log.Lmicroseconds)
	logger.logger[Debug] = log.New(&wr, "[DEBUG] ", log.Ldate|log.Lmicroseconds)

	return &logger
}

func (l *Logger) SetLevel(level int) {
	l.LogLevel = level
	return
}

func (l *Logger) Printf(level int, format string, args ...interface{}) {
	if level < l.LogLevel {
		return
	}

	logger := l.logger[level]
	if logger != nil {
		logger.Printf(format, args...)
	}
}

func (l *Logger) Println(level int, args ...interface{}) {
	if level < l.LogLevel {
		return
	}

	logger := l.logger[level]
	if logger != nil {
		logger.Println(args...)
	}
}
