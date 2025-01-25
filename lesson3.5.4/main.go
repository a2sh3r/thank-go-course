package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	cancelChan := make(chan struct{})
	doneChan := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(dur)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case doneChan <- struct{}{}:
					go func() {
						fn()
						<-doneChan
					}()
				default:
				}
			case <-cancelChan:
				return
			}
		}
	}()

	return func() {
		select {
		case <-cancelChan:
			return
		default:
			close(cancelChan)
		}
	}
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
