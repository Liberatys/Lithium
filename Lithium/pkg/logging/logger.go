package logger

import (
	"log"
)

type Logger interface {
	FormatLog(string) string
	WriteLog(string)
}

type ConsoleLogger struct {
}

func (consoleLogger ConsoleLogger) FormatLog(logMessage string) string {
	return logMessage
}

func (consoleLogger ConsoleLogger) WriteLog(logMessage string) {
	log.Println(logMessage)
}

type FileLogger struct {
	FileLocation string
}

func (fileLogger *FileLogger) WriteLog(string) {

}

func (fileLogger *FileLogger) FormatLog(logMessage string) string {
	return log.Prefix() + "" + logMessage
}
