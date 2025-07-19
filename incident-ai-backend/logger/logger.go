package logger

import (
	"log"
	"os"
)

var Logger *IncidentAILogger

type IncidentAILogger struct {
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
}

func NewLogger() *IncidentAILogger {
	return &IncidentAILogger{
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func InitLogger() {
	Logger = NewLogger()
}

func (l *IncidentAILogger) Info(message string, args ...interface{}) {
	l.infoLog.Printf(message, args...)
}

func (l *IncidentAILogger) Warn(message string, args ...interface{}) {
	l.warnLog.Printf(message, args...)
}

func (l *IncidentAILogger) Error(message string, args ...interface{}) {
	l.errorLog.Printf(message, args...)
}
