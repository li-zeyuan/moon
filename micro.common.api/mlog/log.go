package mlog

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	DeBug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger() *Logger {
	stdout := os.Stderr
	return &Logger{
		DeBug: log.New(io.MultiWriter(stdout), "【debug】", log.LstdFlags|log.Llongfile),
		Info:  log.New(io.MultiWriter(stdout), "【info】", log.LstdFlags|log.Llongfile),
		Error: log.New(io.MultiWriter(stdout), "【error】", log.LstdFlags|log.Llongfile),
	}
}
