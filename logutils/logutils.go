package logutils

import (
	"log"

	"github.com/longbozhan/timewriter"
)

const (
	Debug = iota
	Info
	Warn
	Error
)

type Option struct {
	Dir        string
	Level      int
	Compress   bool
	ReserveDay int
}

type Logger struct {
	printer map[int]*log.Logger
	option  Option
}

func NewLogger(option *Option) *Logger {
	var logger Logger

	wr := timewriter.TimeWriter{
		Dir:        option.Dir,        // 日志路径
		Compress:   option.Compress,   // 是否压缩
		ReserveDay: option.ReserveDay, // 老化天数
	}

	logger.printer = make(map[int]*log.Logger, 0)
	logger.printer[Debug] = log.New(&wr, "[DEBUG] ", log.Ldate|log.Lmicroseconds)
	logger.printer[Info] = log.New(&wr, "[INFO] ", log.Ldate|log.Lmicroseconds)
	logger.printer[Warn] = log.New(&wr, "[WARN] ", log.Ldate|log.Lmicroseconds)
	logger.printer[Error] = log.New(&wr, "[ERROR] ", log.Ldate|log.Lmicroseconds)

	logger.option = *option
	return &logger
}

func (l *Logger) SetLevel(level int) {
	l.option.Level = level
	return
}

func (l *Logger) Printf(level int, format string, args ...interface{}) {
	if level < l.option.Level || level > Error {
		return
	}

	printer := l.printer[level]
	printer.Printf(format, args...)
}

func (l *Logger) Println(level int, args ...interface{}) {
	if level < l.option.Level || level > Error {
		return
	}

	printer := l.printer[level]
	printer.Println(args...)
}
