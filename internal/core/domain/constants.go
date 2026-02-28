// Package domain contains the enterprise business rules and constants
// that are shared across all layers of the application.
package domain

// Application layer labels used for structured logging.
const (
	LogLevelRepository = "Repository"
	LogLevelService    = "Service"
	LogLevelHandler    = "Handler"
	LogLevelMiddleware = "Middleware"
)
