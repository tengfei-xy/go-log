package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var level int

func init() {
	SetLevelInfo()
}
func log(level string, formating string, args ...interface{}) {
	filename, line := "???", 0
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		filename = filepath.Base(filename)
	}
	fmt.Printf("%s %s %s:%d %s\n", level, time.Now().Format("2006-01-02 15:04:05"), filename, line, fmt.Sprintf(formating, args...))
}

const fatal = 0

func SetLevelFatal() {
	level = 0
}
func SetLevelError() {
	level = 1
}
func SetLevelWarn() {
	level = 2
}
func SetLevelInfo() {
	level = 3
}
func SetLevelDebug() {
	level = 4
}
func SetLevelDebug1() {
	level = 5
}
func SetLevelDebug2() {
	level = 6
}
func SetLevelDebug3() {
	level = 7
}
func GetLevel() (int, string) {
	var str string = "Info"

	switch level {
	case 0:
		str = "Fatal"
	case 1:
		str = "Error"
	case 2:
		str = "Warn"
	case 3:
		str = "Info"
	case 4:
		str = "Debug"
	case 5:
		str = "Debug1"
	case 6:
		str = "Debug2"
	case 7:
		str = "Debug3"
	}
	return level, str
}
func Fatalf(formating string, args ...interface{}) {
	if level >= 0 {
		log("F", formating, args...)
		os.Exit(125)
	}
}

func Fatal(formating interface{}) {
	if level >= 0 {
		log("F", fmt.Sprintf("%v", formating))
		os.Exit(125)
	}
}

func Errorf(formating string, args ...interface{}) {
	if level >= 1 {
		log("E", formating, args...)
	}
}
func Error(formating interface{}) {
	if level >= 1 {
		log("E", fmt.Sprintf("%v", formating))
	}
}
func Warnf(formating string, args ...interface{}) {
	if level >= 2 {
		log("W", formating, args...)
	}
}
func Warn(formating interface{}) {
	if level >= 2 {
		log("W", fmt.Sprintf("%v", formating))
	}
}
func Infof(formating string, args ...interface{}) {
	if level >= 3 {
		log("I", formating, args...)
	}
}
func Info(formating interface{}) {
	if level >= 3 {
		log("I", fmt.Sprintf("%v", formating))
	}
}

func Debugf(formating string, args ...interface{}) {
	if level >= 4 {
		log("D", formating, args...)
	}
}
func Debug(formating interface{}) {
	if level >= 4 {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug1f(formating string, args ...interface{}) {
	if level >= 5 {
		log("D", formating, args...)
	}
}
func Debug1(formating interface{}) {
	if level >= 5 {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug2f(formating string, args ...interface{}) {
	if level >= 6 {
		log("D", formating, args...)
	}
}
func Debug2(formating interface{}) {
	if level >= 6 {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug3f(formating string, args ...interface{}) {
	if level >= 7 {
		log("D", formating, args...)
	}
}
func Debug3(formating interface{}) {
	if level >= 7 {
		log("D", fmt.Sprintf("%v", formating))
	}
}
