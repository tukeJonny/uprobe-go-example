package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

type Value struct {
	Name  string
	Value int
}

func sendIntsForever() <-chan *Value {
	stream := make(chan *Value)
	go func() {
		defer close(stream)
		for i := 0; ; i++ {
			if i%2 == 0 {
				continue
			}
			v := &Value{
				Name:  fmt.Sprintf("name%d", i),
				Value: i,
			}
			stream <- v
		}
	}()
	return stream
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	intStream := sendIntsForever()
receiveLoop:
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-intStream:
			fmt.Println(s.Name)
			continue receiveLoop
		}
	}
}
