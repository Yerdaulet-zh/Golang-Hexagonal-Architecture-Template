package kafka

import (
	"fmt"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

// writerLogger implements the segmentio/kafka-go Logger interface
type writerLogger struct {
	core    ports.Logger
	isError bool
}

func (w *writerLogger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if w.isError {
		w.core.Error(msg, "component", "kafka-writer")
	} else {
		w.core.Debug(msg, "component", "kafka-writer")
	}
}
