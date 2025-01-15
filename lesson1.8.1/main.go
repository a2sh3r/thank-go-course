package main

// начало решения

// Produce возвращает срез из n значений val.
func Produce[T any](val T, n int) []T {
	vals := make([]T, n)
	for i := range n {
		vals[i] = val
	}
	return vals
}

// конец решения
