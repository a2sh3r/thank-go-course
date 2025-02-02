package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue chan int

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if !block {
		select {
		case num := <-q:
			return num, nil
		default:
			return 0, ErrEmpty
		}
	}

	num := <-q
	return num, nil
}

// Put помещает элемент в очередь.
// Если очередь заполнена и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if !block {
		select {
		case q <- val:
			return nil
		default:
			return ErrFull
		}
	}

	q <- val
	return nil
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return make(chan int, n)
}

// конец решения

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
