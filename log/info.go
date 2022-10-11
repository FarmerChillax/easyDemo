package log

import (
	"fmt"
	"log"
)

const (
	DEBUG = "[DEBUG]"
	INFO  = "[INFO]"
	WARN  = "[WARN]"
	ERR   = "[ERROR]"
)

func Info(v ...interface{}) {
	log.Println(green(INFO), v)
}

func Infof(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s %s", green(INFO), format), v...)
}

// -------- 日志美化 ----------------------------------------
const (
	COLOR_RED = uint8(iota + 91)
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
)

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_RED, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_GREEN, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_YELLOW, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_BLUE, s)
}
