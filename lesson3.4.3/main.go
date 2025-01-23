package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case out <- randomWord(5):
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			if hasUniqueLetters(word) {
				select {
				case <-cancel:
					return
				case out <- word:
				}
			}
		}
	}()
	return out
}

func hasUniqueLetters(word string) bool {
	charMap := make(map[rune]bool)

	for _, char := range word {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			reversed := reverseString(word)
			select {
			case <-cancel:
				return
			case out <- reversed:
			}
		}
	}()
	return out
}

func reverseString(word string) string {
	a := []rune(word)
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
	return string(a)
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(2)

	out := make(chan string)

	go func() {
		defer wg.Done()
		defer close(out)
		for word := range c1 {
			select {
			case <-cancel:
				return
			case out <- word:
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(out)
		for word := range c2 {
			select {
			case <-cancel:
				return
			case out <- word:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-cancel:
			return
		case word, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(word, " -> ", reverseString(word))
		}
	}
}

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
