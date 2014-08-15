package log

import "os"
import "github.com/cam72cam/go-lumberjack/color"
import "github.com/cam72cam/go-lumberjack/log/logger"

var (
	Debug = logger.New(os.Stdout, color.WhiteThin, "")
	Info = logger.New(os.Stdout, color.WhiteBold, "")
	Warn = logger.New(os.Stdout, color.Yellow, "Warning: ")
	Error = logger.New(os.Stderr, color.Red, "Error: ")
)

func init() {
	SetLevel(level)
}
