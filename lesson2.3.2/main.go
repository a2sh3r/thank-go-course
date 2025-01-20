package main

// не удаляйте импорты, они используются при проверке
import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
type IntSet struct {
	elems map[int]int
}

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	return IntSet{elems: make(map[int]int)}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
	_, ok := s.elems[elem]
	return ok
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s *IntSet) Add(elem int) bool {
	_, ok := s.elems[elem]
	if ok {
		return false
	}
	s.elems[elem] = elem
	return true
}
