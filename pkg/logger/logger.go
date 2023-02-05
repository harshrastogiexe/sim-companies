package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const LFlags = log.LstdFlags | log.Lmicroseconds | log.Lshortfile

var (
	// log level mapped to various prefix
	prefix = map[LogLevel]string{
		LogLevel(Info): "Info",
		LogLevel(Warn): "Warn",
		LogLevel(Fail): "Fail",
	}
	// log level mapped to various log colors
	logColor = map[LogLevel]color.Color{
		LogLevel(Info): *color.New(color.FgGreen, color.Bold),
		LogLevel(Warn): *color.New(color.FgYellow, color.Bold),
		LogLevel(Fail): *color.New(color.BgRed, color.Bold),
	}
	logInstance = log.New(os.Stdout, "", LFlags)
)

func Log(level LogLevel, format string, others ...any) {
	prefixColor := logColor[level]
	logInstance.SetPrefix(prefixColor.Sprintf(prefix[level]) + color.BlackString(" > "))
	logInstance.Output(2, fmt.Sprintf(format, others...))
}
