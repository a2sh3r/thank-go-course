// начало решения

// ZipMap возвращает карту, где ключи - элементы из keys, а значения - из vals.
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {

	m := make(map[K]V)
	minLen := len(vals)
	if len(keys) < len(vals) {
		minLen = len(keys)
	}

	for i := range minLen {
		m[keys[i]] = vals[i]
	}
	return m
}

// конец решения
