package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения
func delay(dur time.Duration, fn func()) func() {
	cancel := make(chan struct{})

	go func() {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case <-timer.C:
			fn()
		case <-cancel:
			return
		}
	}()

	return func() {
		select {
		case <-cancel:
			return
		default:
			close(cancel)
		}
	}
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(130 * time.Millisecond)
}
