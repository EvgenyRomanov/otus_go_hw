package logger

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type logLevel struct {
	label string
	value int
}

var logLevels = map[string]logLevel{
	"error": {
		label: "ERROR",
		value: 4,
	},
	"warning": {
		label: "WARNING",
		value: 3,
	},
	"info": {
		label: "INFO",
		value: 2,
	},
	"debug": {
		label: "DEBUG",
		value: 1,
	},
}

type Logger struct {
	level   logLevel
	writeTo io.Writer
}

func New(level string, writeTo io.Writer) *Logger {
	level = strings.TrimSpace(strings.ToLower(level))

	targetLvl, found := logLevels[level]
	if !found {
		targetLvl = logLevels["debug"]
	}

	return &Logger{targetLvl, writeTo}
}

func (l Logger) core(level logLevel, msg string, params ...any) {
	// do not write anything if request level less then in config
	if level.value < l.level.value {
		return
	}

	var logString strings.Builder
	logString.WriteString(fmt.Sprintf("%s [%s] ", time.Now().Format("2006-01-02 15:04:05"), level.label))

	if params != nil {
		logString.WriteString(fmt.Sprintf(msg, params...))
	} else {
		logString.WriteString(msg)
	}

	logString.WriteString("\n")
	l.writeTo.Write([]byte(logString.String()))
}

func (l Logger) Info(msg string, params ...any) {
	l.core(logLevels["info"], msg, params...)
}

func (l Logger) Error(msg string, params ...any) {
	l.core(logLevels["error"], msg, params...)
}

func (l Logger) Warning(msg string, params ...any) {
	l.core(logLevels["warning"], msg, params...)
}

func (l Logger) Debug(msg string, params ...any) {
	l.core(logLevels["debug"], msg, params...)
}

func (l Logger) Log(msg string, params ...any) {
	l.core(l.level, msg, params...)
}
