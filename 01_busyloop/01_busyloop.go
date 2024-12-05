package main

import (
	"context"
	"crypto/rand"
	"math/big"
	"os/signal"
	"syscall"
)

func doSomething() {
	for i := 0; i < 10; i++ {
		rand.Int(rand.Reader, big.NewInt(100))
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			doSomething()
		}
	}
}
