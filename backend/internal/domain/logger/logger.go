package logger

type Logger interface {
	LogError(code, backendMessage string, details map[string]string, err error)
}
