// начало решения

// Avg - накопительное среднее значение.
type Avg[T int | float64 ] struct { 
    val T
    count int
}

// Add пересчитывает среднее значение с учетом val.
func (a *Avg[T]) Add(val T) *Avg[T] {
    a.val = a.val + T(val)
    a.count += 1
    return a
}

// Val возвращает текущее среднее значение.
func (a Avg[T]) Val() T {
    res := T(0)
    if a.count != 0 {
        res = a.val/T(a.count)
    }
    return res
}

// конец решения
