package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
    var res []rune
    
	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
    if len(text) > width {
        res = (append([]rune(text)[:width], '.','.','.'))
    } else if len(text) <= width {
        res = []rune(text)
    }
    
    fmt.Println(string(res))
}
