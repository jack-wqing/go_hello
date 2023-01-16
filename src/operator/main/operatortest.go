package main

import "fmt"

func main() {

	//算数运算符 数值运算
	fmt.Println(10 / 4)

	var n2 float32 = 10.0 / 4
	fmt.Println("n2 = ", n2)

	//++ -- 操作符只能独立使用
	n2++
	fmt.Printf("n2++ %v \n", n2)
	n2--
	n2--
	fmt.Printf("n-- %v \n", n2)

	//关系运算符
	var num1 int = 12
	var num2 int = 13
	fmt.Println("num1 > num2 = ", num1 > num2)

	//逻辑操作符
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("age > 30 && age < 50")
	}
	if age > 20 && age < 40 {
		fmt.Println("age > 30 && age < 40")
	}

	//赋值运算符
	var name string = "小乔"
	var nickName string = "别名:"
	nickName += name
	fmt.Println("nickName:", nickName)

	//其他操作符
	var namePtr *string = &name
	fmt.Printf("name &= %d, value= %v", &name, *namePtr)

	//三元操作的方式 go中是没有 ? : 三元操作符的
	var nValue int
	var iValue int = 1
	var jValue int = 13
	if iValue > 2 {
		nValue = iValue
	} else {
		nValue = jValue
	}
	fmt.Println("nValue = ", nValue)

}
