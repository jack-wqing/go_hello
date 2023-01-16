package main

import "fmt"

// init函数会被 Go框架先 与 main 函数执行
// init 通常可以执行一些初始化的工作
func init() {
	fmt.Println("init func test ....")
}

func sumAndSub(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func test(n int) {
	if n >= 2 {
		n--
		test(n)
	}
	fmt.Printf("test n=%d \n", n)
}

func test1(n int) {
	if n >= 2 {
		n--
		test1(n)
	} else {
		fmt.Println("test1 n=", n)
	}
}

func pointTest(value *int) {
	*value = *value * 2
}

// 计算斐波拉契数
func fbn(n int) int {
	if n == 1 || n == 2 {
		//fmt.Printf("fbn n=%v, value=%v \n", n, 1)
		return 1
	} else {
		result := fbn(n-1) + fbn(n-2)
		//fmt.Printf("fbn n=%v, value=%v \n", n, result)
		return result
	}
}

func funTestSum(n1 int, n2 int) int {
	return n1 + n2
}

// 给函数起个别名，方便使用, 定义放在使用的前面
type myFunType func(int, int) int

func myFun(fun myFunType, n1 int, n2 int) int {
	return fun(n1, n2)
}

// 给函数类型定义别名
type myReturnNameType func(int, int) (int, int)

// 函数返回值也可以指定返回的名字
func myReturnName(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func allSum(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// go不支持传统的 函数重载
func main() {

	num1 := 12
	num2 := 13

	sum, sub := sumAndSub(num1, num2)
	fmt.Printf("sum : %v, sum : %v \n", sum, sub)
	// := 会进行变量推导 使用 "_" 占位时，直接使用 = 赋值
	sum, _ = sumAndSub(num1, num2)
	fmt.Println("sum=", sum)

	test(4)
	test1(4)
	//斐波拉契数
	fbnRes := fbn(6)
	fmt.Println("fbnRes:", fbnRes)

	pointInt := 1
	fmt.Println("pointInt before:", pointInt)
	pointTest(&pointInt)
	fmt.Println("pointInt after:", pointInt)

	//函数作为数据类型的操作
	funTestVar := funTestSum
	fmt.Printf("funcTestVar: %T, funTestSum: %T \n", funTestVar, funTestSum)
	funTestSumValue := funTestVar(1, 2)
	fmt.Println("fumTestSumValue:", funTestSumValue)
	//函数可以作为形参使用
	fmt.Printf("myFun Value: %d", myFun(funTestSum, 12, 23))

	//可以通过type 关键字给数据类型起别名 使用的时候需要显示转换
	type myInt int
	var myIntValue myInt
	var originIntValue int
	myIntValue = 12
	originIntValue = 13
	fmt.Printf("myIntValue: %v, originIntValue: %v \n", myIntValue, originIntValue)

	mySum, mySub := myReturnName(1, 2)
	fmt.Printf("mySum: %v, mySub: %v \n", mySum, mySub)

	//可变参数
	allSumRes := allSum(1, 2, 3)
	fmt.Printf("allSum:%d \n", allSumRes)
}
