package system

import (
	"log"
	"sync"
)

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

var logLevelNames = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

type Logger struct {
	level LogLevel
}

var (
	instance *Logger
	once     sync.Once
)

func (logger *Logger) log(msg string, level LogLevel) {

	if logger.level <= level {
		var loglevelName = "Unknown"
		if int(level) < len(logLevelNames) {
			loglevelName = logLevelNames[level]
		}
		log.Printf("[%s] %s", loglevelName, msg)
	}
}

func (logger *Logger) Trace(msg string) {
	logger.log(msg, TRACE)
}

func (logger *Logger) Debug(msg string) {
	logger.log(msg, DEBUG)
}

func (logger *Logger) Info(msg string) {
	logger.log(msg, INFO)
}

func (logger *Logger) Warn(msg string) {
	logger.log(msg, WARN)
}

func (logger *Logger) Error(msg string) {
	logger.log(msg, ERROR)
}

func (g *Logger) Fatal(msg string) {
	log.Fatalf(msg)
}

func GetLogger(level LogLevel) *Logger {
	once.Do(func() {
		instance = &Logger{level: level}
	})
	return instance
}
