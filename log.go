package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var LOGLEVEL int

func log(level string, formating string, args ...interface{}) {
	filename, line := "???", 0
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		filename = filepath.Base(filename)
	}
	fmt.Printf("%s %s %s:%d %s\n", level, time.Now().Format("2006-01-02 15:04:05"), filename, line, fmt.Sprintf(formating, args...))

}

func Debugf(formating string, args ...interface{}) {
	if LOGLEVEL <= 0 {
		log("D", formating, args...)
	}
}
func Debug(formating interface{}) {
	if LOGLEVEL <= 0 {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Infof(formating string, args ...interface{}) {
	if LOGLEVEL <= 1 {
		log("I", formating, args...)
	}
}
func Info(formating interface{}) {
	if LOGLEVEL <= 1 {
		log("I", fmt.Sprintf("%v", formating))
	}
}
func Warnf(formating string, args ...interface{}) {
	if LOGLEVEL <= 2 {
		log("W", formating, args...)
	}
}
func Warn(formating interface{}) {
	if LOGLEVEL <= 2 {
		log("W", fmt.Sprintf("%v", formating))
	}
}

func Errorf(formating string, args ...interface{}) {
	if LOGLEVEL <= 3 {
		log("E", formating, args...)
	}
}
func Error(formating interface{}) {
	if LOGLEVEL <= 3 {
		log("E", fmt.Sprintf("%v", formating))
	}
}
func Fatalf(formating string, args ...interface{}) {
	if LOGLEVEL <= 4 {
		log("F", formating, args...)
		os.Exit(125)
	}
}
func Fatal(formating interface{}) {
	if LOGLEVEL <= 4 {
		log("F", fmt.Sprintf("%v", formating))
		os.Exit(125)
	}
}
