package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"
)

func appendMany(ctx context.Context) {
	s := []any{}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s = append(s, struct{}{})
		}
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go appendMany(ctx)

	<-ctx.Done()
}
