package logging

import (
	"log"
	"os"
	"sync"
)

// Logger is the structure that holds the logger instance.
type Logger struct {
	logger *log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

// GetLogger returns the singleton logger instance.
func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			// Use os.Stdout for logging, which implements io.Writer.
			logger: log.New(os.Stdout, "", log.LstdFlags),
		}
	})
	return instance
}

// LogWithCorrelationID logs a message along with the correlation ID.
func (l *Logger) LogWithCorrelationID(correlationID, message string) {
	l.logger.Printf("[CorrelationID: %s] %s", correlationID, message)
}
