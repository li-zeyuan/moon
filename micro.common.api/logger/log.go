package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	PanicLevel = 1
	FatalLevel = 2
	ErrorLevel = 3
	WarnLevel  = 4
	InfoLevel  = 5
	DebugLevel = 6
)

const (
	prefixPanic = "[PANIC] "
	prefixFatal = "[FATAL] "
	PrefixError = "[ERROR] "
	prefixWarn  = "[WARN] "
	prefixInfo  = "[INFO] "
	prefixDebug = "[DEBUG] "
)

var DefaultLogger *Logger

type Logger struct {
	mLog      *log.Logger
	RequestId string
}

func init() {
	DefaultLogger = NewLogger("")
}

func NewLogger(requestId string) *Logger {
	stdout := os.Stderr
	return &Logger{
		mLog:      log.New(io.MultiWriter(stdout), "", log.LstdFlags|log.Llongfile|log.Lmsgprefix),
		RequestId: requestId,
	}
}

func (l *Logger) formalRequestId() string {
	if len(l.RequestId) == 0 {
		return ""
	}
	return fmt.Sprintf(" {request_id: %s}", l.RequestId)
}

func (l *Logger) output(calldepth int, s string) {
	s = s + l.formalRequestId()
	l.mLog.Output(calldepth, s)
}

func (l *Logger) Debug(v ...interface{}) {
	l.mLog.SetPrefix(blue(prefixDebug))
	l.output(2, fmt.Sprint(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.mLog.SetPrefix(green(prefixInfo))
	l.output(3, fmt.Sprint(v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.mLog.SetPrefix(yellow(prefixWarn))
	l.output(3, fmt.Sprint(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.mLog.SetPrefix(red(PrefixError))
	l.output(3, fmt.Sprint(v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.mLog.SetPrefix(magenta(prefixFatal))
	l.output(3, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	l.mLog.SetPrefix(red(prefixPanic))
	l.output(3, fmt.Sprint(v...))
	panic(v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.mLog.SetPrefix(blue(prefixDebug))
	l.output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.mLog.SetPrefix(green(prefixInfo))
	l.output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.mLog.SetPrefix(yellow(prefixWarn))
	l.output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.mLog.SetPrefix(red(PrefixError))
	l.output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.mLog.SetPrefix(magenta(prefixFatal))
	l.output(3, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.mLog.SetPrefix(red(prefixPanic))
	l.output(3, fmt.Sprintf(format, v...))
	panic(v)
}
