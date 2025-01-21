package main

// не удаляйте импорты, они используются при проверке
import (
	"strings"
)

// Words работает со словами в строке.
type Words struct {
	words []string
}

// MakeWords создает новый экземпляр Words.
func MakeWords(s string) Words {
	return Words{words: strings.Split(s, " ")}
}

// Index возвращает индекс первого вхождения слова в строке,
// или -1, если слово не найдено.
func (w Words) Index(word string) int {
	for idx, item := range w.words {
		if item == word {
			return idx
		}
	}
	return -1
}
