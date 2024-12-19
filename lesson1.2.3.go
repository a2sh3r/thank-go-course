package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
    var result string
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
    // ...
    
    for _ = range times {
        result += source
    }

	fmt.Println(result)
}
