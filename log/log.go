package log

import (
	"io"
	slog "log"
	"os"
)

// ERROR : Err logger
var ERROR *slog.Logger

// WARN : Warning logger
var WARN *slog.Logger

// INFO : info logger
var INFO *slog.Logger

var f *os.File

// InitLogger : Init Global log
func InitLogger(path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		slog.Fatalln("Create log file failed")
	}

	output := io.MultiWriter(os.Stdout, f)

	ERROR = slog.New(output, "[ERR]", slog.Ldate|slog.Ltime|slog.Lshortfile)
	WARN = slog.New(output, "[WARN]", slog.Ldate|slog.Ltime|slog.Lshortfile)
	INFO = slog.New(output, "[INFO]", slog.Ldate|slog.Ltime)
}

// FiniLogger : Close Log File Handdler
func FiniLogger() {
	f.Close()
}
