package logger

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

var Log Logger = Logger{}

type LoggerOptions struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type LoggerSetUpOptions struct {
	Info  *LoggerOptions
	Warn  *LoggerOptions
	Error *LoggerOptions
}

func Init(options *LoggerSetUpOptions) error {
	if options.Info == nil {
		return fmt.Errorf("info options are missings")
	}

	if options.Warn == nil {
		return fmt.Errorf("warn options are missings")
	}

	if options.Error == nil {
		return fmt.Errorf("error options are missings")
	}

	Log.Info = CreateLogger(options.Info, log.InfoLevel)
	Log.Warn = CreateLogger(options.Warn, log.WarnLevel)
	Log.Error = CreateLogger(options.Error, log.ErrorLevel)
	return nil
}

func CreateLogger(options *LoggerOptions, level log.Level) *log.Logger {
	new_log := log.New()
	new_log.SetLevel(level)
	new_log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	new_log.SetOutput(&lumberjack.Logger{
		Filename:   options.Filename,
		MaxSize:    options.MaxAge,
		MaxBackups: options.MaxBackups,
		MaxAge:     options.MaxAge,
		Compress:   options.Compress,
	})
	return new_log

}
