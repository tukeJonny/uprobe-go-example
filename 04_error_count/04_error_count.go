package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os/signal"
	"syscall"
	"time"
)

// fmt.Errorf呼び出し箇所で、スタックとあわせて出力
// 呼び出し箇所ごとカウント出せると良いかも
// @[stack] = count();

func f1() error {
	return fmt.Errorf("error f1")
}

func f2() error {
	return fmt.Errorf("error f2")
}

func f3() error {
	return fmt.Errorf("error f3")
}

func gen() error {
	prob, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return fmt.Errorf("failed to generate probability: %w", err)
	}

	if prob.Int64() < 30 {
		if err := f1(); err != nil {
			return err
		}
	} else if prob.Int64() < 60 {
		if err := f2(); err != nil {
			return err
		}
	} else {
		if err := f3(); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			gen()
		}
	}
}
