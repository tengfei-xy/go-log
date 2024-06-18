package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

const LEVELFATAL string = "fatal"
const LEVELERROR string = "error"
const LEVELWARN string = "warn"
const LEVELINFO string = "info"
const LEVELDEBUG string = "debug"
const LEVELDEBUG1 string = "debug1"
const LEVELDEBUG2 string = "debug2"
const LEVELDEBUG3 string = "debug3"

const LEVELFATALINT int = 0
const LEVELERRORINT int = 1
const LEVELWARNINT int = 2
const LEVELINFOINT int = 3
const LEVELDEBUGINT int = 4
const LEVELDEBUG1INT int = 5
const LEVELDEBUG2INT int = 6
const LEVELDEBUG3INT int = 7

func SetLevelStr(v string) {
	switch strings.ToLower(v) {
	case LEVELFATAL:
		SetLevelFatal()
	case LEVELERROR:
		SetLevelError()
	case LEVELWARN:
		SetLevelError()
	case LEVELINFO:
		SetLevelInfo()
	case LEVELDEBUG:
		SetLevelDebug()
	case LEVELDEBUG1:
		SetLevelDebug1()
	case LEVELDEBUG2:
		SetLevelDebug2()
	case LEVELDEBUG3:
		SetLevelDebug3()
	default:
		SetLevelInfo()
	}
}
func SetLevelInt(v int) {
	switch v {
	case LEVELFATALINT:
		SetLevelFatal()
	case LEVELERRORINT:
		SetLevelError()
	case LEVELWARNINT:
		SetLevelError()
	case LEVELINFOINT:
		SetLevelInfo()
	case LEVELDEBUGINT:
		SetLevelDebug()
	case LEVELDEBUG1INT:
		SetLevelDebug1()
	case LEVELDEBUG2INT:
		SetLevelDebug2()
	case LEVELDEBUG3INT:
		SetLevelDebug3()
	default:
		SetLevelInfo()
	}
}
func SetLevelFatal() (int, string) {
	level = LEVELFATALINT
	return level, "Fatal"

}
func SetLevelError() (int, string) {
	level = LEVELERRORINT
	return level, "Error"

}
func SetLevelWarn() (int, string) {
	level = LEVELWARNINT
	return level, "Warn"

}
func SetLevelInfo() (int, string) {
	level = LEVELINFOINT
	return level, "Info"

}
func SetLevelDebug() (int, string) {
	level = LEVELDEBUGINT
	return level, "Debug"

}
func SetLevelDebug1() (int, string) {
	level = LEVELDEBUG1INT
	return level, "Debug1"

}
func SetLevelDebug2() (int, string) {
	level = LEVELDEBUG2INT
	return level, "Debug2"

}
func SetLevelDebug3() (int, string) {
	level = LEVELDEBUG3INT
	return level, "Debug3"
}
func GetLevel() (int, string) {
	var str string

	switch level {
	case LEVELFATALINT:
		str = LEVELFATAL
	case LEVELERRORINT:
		str = LEVELERROR
	case LEVELWARNINT:
		str = LEVELWARN
	case LEVELINFOINT:
		str = LEVELINFO
	case LEVELDEBUGINT:
		str = LEVELDEBUG
	case LEVELDEBUG1INT:
		str = LEVELDEBUG1
	case LEVELDEBUG2INT:
		str = LEVELDEBUG2
	case LEVELDEBUG3INT:
		str = LEVELDEBUG3
	default:
		str = LEVELINFO
	}
	return level, str
}
func Fatalf(formating string, args ...interface{}) {
	if level >= LEVELFATALINT {
		log("F", formating, args...)
		os.Exit(125)
	}
}

func Fatal(formating interface{}) {
	if level >= LEVELFATALINT {
		log("F", fmt.Sprintf("%v", formating))
		os.Exit(125)
	}
}

func Errorf(formating string, args ...interface{}) {
	if level >= LEVELERRORINT {
		log("E", formating, args...)
	}
}
func Error(formating interface{}) {
	if level >= LEVELERRORINT {
		log("E", fmt.Sprintf("%v", formating))
	}
}
func Warnf(formating string, args ...interface{}) {
	if level >= LEVELWARNINT {
		log("W", formating, args...)
	}
}
func Warn(formating interface{}) {
	if level >= LEVELWARNINT {
		log("W", fmt.Sprintf("%v", formating))
	}
}
func Infof(formating string, args ...interface{}) {
	if level >= LEVELINFOINT {
		log("I", formating, args...)
	}
}
func Info(formating interface{}) {
	if level >= LEVELINFOINT {
		log("I", fmt.Sprintf("%v", formating))
	}
}

func Debugf(formating string, args ...interface{}) {
	if level >= LEVELDEBUGINT {
		log("D", formating, args...)
	}
}
func Debug(formating interface{}) {
	if level >= LEVELDEBUGINT {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug1f(formating string, args ...interface{}) {
	if level >= LEVELDEBUG1INT {
		log("D", formating, args...)
	}
}
func Debug1(formating interface{}) {
	if level >= LEVELDEBUG1INT {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug2f(formating string, args ...interface{}) {
	if level >= LEVELDEBUG2INT {
		log("D", formating, args...)
	}
}
func Debug2(formating interface{}) {
	if level >= LEVELDEBUG2INT {
		log("D", fmt.Sprintf("%v", formating))
	}
}
func Debug3f(formating string, args ...interface{}) {
	if level >= LEVELDEBUG3INT {
		log("D", formating, args...)
	}
}
func Debug3(formating interface{}) {
	if level >= LEVELDEBUG3INT {
		log("D", fmt.Sprintf("%v", formating))
	}
}
