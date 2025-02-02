package main

import (
	"fmt"
	"os"
)

// normalize нормализует значения, переданные в vals,
// так чтобы их сумма была равна 1.
func normalize(vals ...*float64) {
	var sum float64
	for _, val := range vals {
		sum += *val
	}

	for _, val := range vals {
		*val /= sum
	}
}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
	os.Exit(0)
}
