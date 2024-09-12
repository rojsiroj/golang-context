package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// go routine leak
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:

				time.Sleep(1 * time.Second)
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestGoRoutineLeak(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	parent := context.Background()

	destination := CreateCounter(parent) // make go routine leak because always send to channel destination
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(5 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}
