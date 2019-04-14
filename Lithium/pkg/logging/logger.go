package logger

type Logger interface {
	FormatLog(string) string
	WriteLog(string)
}
