package log

import (
	"fmt"
	"os"
)

// Logger defines the operations required to support logging within the
// application.
type Logger interface {
	Info(a ...interface{})
}

// DefaultLogger logs to stderr for Error messages and to stdout for
// Info messages.
type DefaultLogger struct{}

// Info logs to stdout
func (l *DefaultLogger) Info(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}
