package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	phrase := "любитель французских булок"

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
    // ...
    var abbr []rune
    
    words := strings.Fields(phrase)
    for _,v := range words {
        fmt.Println(string(v))
    }
    
    for _,v := range strings.Fields(phrase) {
        if len(v) > 0 && unicode.IsLetter(rune(v[0])) {
            abbr = append(abbr, ([]rune(v)[0]))
        }
    }
    
	fmt.Println(string(abbr))
}


