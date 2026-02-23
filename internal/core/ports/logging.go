// Package ports defines the abstractions for driving and driven adapters.
package ports

// Logger defines the contract for logging across the application.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}
