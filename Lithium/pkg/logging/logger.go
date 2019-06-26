package logger

import (
	"fmt"
	"log"
)

type Logger interface {
	FormatLog(string) string
	WriteLog(string)
}

type ConsoleLogger struct {
}

func (consoleLogger ConsoleLogger) FormatLog(logMessage string) string {
	return log.Prefix() + "" + logMessage
}

func (consoleLogger ConsoleLogger) WriteLog(logMessage string) {
	fmt.Println(logMessage)
}

type FileLogger struct {
	FileLocation string
}

func (fileLogger *FileLogger) WriteLog(string) {

}

func (fileLogger *FileLogger) FormatLog(logMessage string) string {
	return log.Prefix() + "" + logMessage
}

type StatisticsLogger struct {
	Request int
}

func (statisticsLogger StatisticsLogger) FormatLog(logMessage string) string {
	return ""
}
func (statisticsLogger StatisticsLogger) WriteLog(logMessage string) {
	return
}
