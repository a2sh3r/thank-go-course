package main

// начало решения

// Map - карта "ключ-значение".
type Map[K comparable, V any] map[K]V

// Set устанавливает значение для ключа.
func (m Map[K, V]) Set(key K, val V) {
	m[key] = val
}

// Get возвращает значение по ключу.
func (m Map[K, V]) Get(key K) V {
	return m[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m Map[K, V]) Keys() []K {
	var keys []K
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m Map[K, V]) Values() []V {
	var vals []V
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

// конец решения
