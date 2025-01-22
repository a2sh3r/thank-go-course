package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say печатает фразу от имени обработчика
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// начало решения

// makePool создает пул на n обработчиков
// возвращает функции handle и wait
func makePool(n int, handler func(int, string)) (func(string), func()) {

	pool := make(chan int, n)
	done := make(chan struct{})

	for i := 0; i < n; i++ {
		pool <- i
	}

	taskCounter := 0

	handle := func(phrase string) {
		id := <-pool
		taskCounter++
		go func() {
			defer func() {
				pool <- id
				taskCounter--
				if taskCounter == 0 {
					close(done)
				}
			}()
			handler(id, phrase)
		}()
	}

	wait := func() {
		if taskCounter > 0 {
			<-done
		}
	}

	return handle, wait
}

// конец решения

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	handle, wait := makePool(2, say)
	for _, phrase := range phrases {
		handle(phrase)
	}
	wait()
}
