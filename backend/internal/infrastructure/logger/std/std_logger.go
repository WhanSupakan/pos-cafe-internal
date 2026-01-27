package std

import (
	"cafe-pos/internal/domain/logger"
	"fmt"
	"log"
)

type StdLogger struct{}

func NewStdLogger() logger.Logger {
	return &StdLogger{}
}

func (l *StdLogger) LogError(code, backendMessage string, details map[string]string, err error) {
	msg := fmt.Sprintf("[%s] %s", code, backendMessage)
	if err != nil {
		msg += fmt.Sprintf(" | Error: %v", err)
	}
	if len(details) > 0 {
		msg += fmt.Sprintf(" | Details: %v", details)
	}
	log.Println(msg)
}
