package service

import "testing"

func TestFactorial(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Factorial(i)
	}
	PrintDataMap()
}
