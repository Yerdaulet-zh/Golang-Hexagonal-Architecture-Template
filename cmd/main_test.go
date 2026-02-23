package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	err := run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		t.Errorf("run() failed: %v", err)
	}
}
