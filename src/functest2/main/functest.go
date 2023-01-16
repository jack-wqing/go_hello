package main

import "fmt"

var (
	allLength = func(params ...int) int { // 这种情况不能使用 var funanme := func(){} 类型推断的方式
		return len(params)
	}
)

func addUpper() func(int) int {
	var n = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// 对于defer的测试
func sum(n1 int, n2 int) int {
	//当执行到defer时， 会将语句压入独立的栈，暂时不执行
	//当函数执行完后，按照先入后出的方式执行
	defer fmt.Println("defer n1=", n1)
	defer fmt.Println("defer n2=", n2)
	res := n1 + n2
	fmt.Println("defer res=", res)
	return res
}

var name = "jack"

func main() {
	//1、定义命名函数，直接调用
	niMingFunc := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Printf("niminFunc: %v \n", niMingFunc)

	subFunc := func(n1, n2 int) int {
		return n1 - n2
	}
	niMingVar := subFunc(1, 3)
	fmt.Printf("niMingVar: %d \n", niMingVar)

	allLengthValue := allLength(1, 2, 3)
	fmt.Printf("allLengthValue: %d \n", allLengthValue)

	//闭包操作：练习操作
	upperBase := addUpper()
	baseValue := upperBase(12)
	fmt.Printf("baseValue: %d \n", baseValue)
	baseValue1 := upperBase(12)
	fmt.Printf("baseValue1: %d \n", baseValue1)
	// defer测试
	sumValue := sum(10, 20)
	fmt.Println("defer main sumValue:", sumValue)

	// 测试变量的同名 全局与局部变量声明重复的名字：使用是就近原则使用
	var name = "jack"
	fmt.Printf("name repeat : %s", name)

}
