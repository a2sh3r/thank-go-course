// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	ticker := time.NewTicker(time.Second / time.Duration(limit))
	cancelChan := make(chan struct{})
	requestChan := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case requestChan <- struct{}{}:
				default:
				}
			case <-cancelChan:
				return
			}
		}
	}()

	handle = func() error {
		select {
		case <-cancelChan:
			return ErrCanceled
		case <-requestChan:
			fn()
			return nil
		}
	}
	cancel = func() {
		select {
		case <-cancelChan:
			return
		default:
			close(cancelChan)
		}
	}

	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(10, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
