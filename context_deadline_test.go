package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithDeadLine(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) //fixed time deadline
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}
