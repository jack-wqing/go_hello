package pactest

import "fmt"

func init() {
	fmt.Println("util init test ...")
}

// Calculate 需要将函数名首字母大写
func Calculate(num1 int, num2 int, operator byte) int {
	var result int
	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	default:
		fmt.Println("操作符操作")
	}
	return result
}
