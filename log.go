package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func log(level string, formating string, args ...interface{}) {
	filename, line := "???", 0
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		filename = filepath.Base(filename)
	}
	fmt.Printf("%s %s:%d %s %s\n", time.Now().Format("2006-01-02 15:04:05"), filename, line, level, fmt.Sprintf(formating, args...))

}

func Debugf(formating string, args ...interface{}) {
	if LOGLEVEL <= 0 {
		log("debug", formating, args...)
	}
}
func Debug(formating interface{}) {
	if LOGLEVEL <= 0 {
		log("debug", fmt.Sprintf("%v", formating))
	}
}
func Infof(formating string, args ...interface{}) {
	if LOGLEVEL <= 1 {
		log("info ", formating, args...)
	}
}
func Info(formating interface{}) {
	if LOGLEVEL <= 1 {
		log("info ", fmt.Sprintf("%v", formating))
	}
}
func Warnf(formating string, args ...interface{}) {
	if LOGLEVEL <= 2 {
		log("warn ", formating, args...)
	}
}
func Warn(formating interface{}) {
	if LOGLEVEL <= 2 {
		log("warn ", fmt.Sprintf("%v", formating))
	}
}

func Errorf(formating string, args ...interface{}) {
	if LOGLEVEL <= 3 {
		log("error", formating, args...)
	}
}
func Error(formating interface{}) {
	if LOGLEVEL <= 3 {
		log("error", fmt.Sprintf("%v", formating))
	}
}
func Fatalf(formating string, args ...interface{}) {
	if LOGLEVEL <= 4 {
		log("fatal", formating, args...)
		os.Exit(125)
	}
}
func Fatal(formating interface{}) {
	if LOGLEVEL <= 4 {
		log("fatal", fmt.Sprintf("%v", formating))
		os.Exit(125)
	}
}
