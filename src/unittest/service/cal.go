package service

func AddUpper(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func Sub(n1, n2 int) int {
	return n1 - n2
}
