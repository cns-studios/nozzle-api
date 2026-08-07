package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// TODO: Auditlogs using an database

var logger *log.Logger

func SetupLogging() {

	var writer io.Writer = os.Stdout

	logger = log.New(writer, "", 0)
}

func logLine(level, format string, args ...any) string {
	msg := fmt.Sprintf(format, args...)
	return fmt.Sprintf("%s [%s] %s", time.Now().Format("2006-01-02 15:04:05"), level, msg)
}

func LogInfo(format string, args ...any) {
	if logger == nil {
		return
	}
	logger.Println(logLine("INFO", format, args...))
}

func LogWarn(format string, args ...any) {
	if logger == nil {
		return
	}
	logger.Println(logLine("WARNING", format, args...))
}

func LogError(format string, args ...any) {
	if logger == nil {
		return
	}
	logger.Println(logLine("ERROR", format, args...))
}
