package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

type lokiAdapter struct {
	labels map[string]string
	url    string
}

func NewLokiLogger(url string, labels map[string]string) ports.Logger {
	return &lokiAdapter{
		url:    url,
		labels: labels,
	}
}

func (l *lokiAdapter) send(level, msg string, args ...any) {
	line := fmt.Sprintf("level=%s msg=%q", level, msg)

	// Loop through args to create individual rows/keys
	// We assume args come in pairs: ["key", value, "key2", value2]
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key := fmt.Sprintf("%v", args[i])
			val := fmt.Sprintf("%v", args[i+1])
			// Append each as its own key=value pair
			line = fmt.Sprintf("%s %s=%q", line, key, val)
		} else {
			// If there's an odd number of args, just append the last one
			line = fmt.Sprintf("%s extra=%q", line, args[i])
		}
	}

	ts := fmt.Sprintf("%d", time.Now().UnixNano())
	lokiMsg := map[string]any{
		"streams": []map[string]any{
			{
				"stream": l.labels,
				"values": [][]string{{ts, line}},
			},
		},
	}

	body, err := json.Marshal(lokiMsg)
	if err != nil {
		fmt.Printf("logging error: could not marshal loki msg: %v\n", err)
		return
	}

	resp, err := http.Post(l.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("logging error: failed to send to loki: %v\n", err)
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error while closing the loki post response body: %v\n", err)
		}
	}()

	if resp.StatusCode >= 400 {
		fmt.Printf("logging error: loki returned status %d\n", resp.StatusCode)
	}
}

func (l *lokiAdapter) Debug(msg string, args ...any) {
	l.send("Debug", msg, args...)
}

func (l *lokiAdapter) Info(msg string, args ...any) {
	l.send("Info", msg, args...)
}

func (l *lokiAdapter) Warn(msg string, args ...any) {
	l.send("Warn", msg, args...)
}

func (l *lokiAdapter) Error(msg string, args ...any) {
	l.send("Error", msg, args...)
}
