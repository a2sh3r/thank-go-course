package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)
    var lang string
    
    switch code {
    case "en":
        lang = "English"
    case "fr":
        lang = "French"
    case "ru", "rus" :
        lang = "Russian"
    default:
        lang = "Unknown"   
    }

	fmt.Println(lang)
}
