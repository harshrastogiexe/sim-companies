package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Info  = log.New(os.Stdout, color.CyanString("[INFO] "), log.Default().Flags())
	Error = log.New(os.Stdout, color.RedString("[ERROR] "), log.Default().Flags())
	Warn  = log.New(os.Stdout, color.RedString("[WARN] "), log.Default().Flags())
)
