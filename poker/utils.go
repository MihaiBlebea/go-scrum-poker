package poker

func fibonacci(n int) []uint {
	a := 1
	b := 1
	list := make([]uint, 0)

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
		list = append(list, uint(a))
	}

	return list
}
