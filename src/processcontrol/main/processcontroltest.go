package main

import (
	"fmt"
	"go/types"
	"math/rand"
	"time"
)

func main() {

	//1，单分支运行
	var i int = 12
	if i < 13 {
		fmt.Println("i < 13")
	}
	//2、双分支
	if i < 12 {
		fmt.Println("i < 12")
	} else {
		fmt.Println("i >= 12")
	}
	//3、多分支
	if i < 12 {
		fmt.Println("i < 12")
	} else if iv := 2; i > 10 {
		fmt.Println("i > 10, iv=", iv)
	} else {
		fmt.Println("default else ")
	}

	// go支持值if语句条件哪里声明一个变量
	if age := 20; age > 12 {
		fmt.Println("age > 12 && age=", age)
	}
	//switch 语句的方式
	var iSwitch int = 3
	switch iSwitch {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	var age int = 20
	//switch后面可以不带表达式
	switch {
	case age == 20:
		fmt.Print("age == 20")
	case age == 10:
		fmt.Println("age == 10")
	default:
		fmt.Println("age default")
	}
	// switch 后面可以定义变量 分号结束
	// 可以使用 fallthrough 接着执行下一个case 不用看条件
	switch address := "aa"; {
	case address == "bb":
		fmt.Println("address == bb")
	}

	var x interface{}
	var y float64 = 10.0
	x = y
	switch i := x.(type) { //x.(type) 可以输出x的变量类型
	case types.Nil:
		fmt.Printf("x的类型～ : %T \n", i)
	case int:
		fmt.Printf("x 是 int 类型 \n")
	case float64:
		fmt.Printf("x 是 float64 类型 \n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 类型 \n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 类型 \n")
	default:
		fmt.Printf("x 是 未知型 \n")
	}

	//循环代码测试
	//for得第一种
	for i := 1; i < 5; i++ {
		fmt.Println("i=", i)
	}
	//for的第二种
	index := 1
	for index < 5 {
		index++
		fmt.Println("index condition:", index)
	}
	//for的第三种
	for {
		index++
		if index >= 10 {
			fmt.Println("index=", index)
			break
		}
	}
	//for range遍历
	strValue := "hello world ！ 中国"
	strValueRune := []rune(strValue)         //数据切片 可以处理字符串中文遍历的问题
	for i := 0; i < len(strValueRune); i++ { //传统的字符串遍历是按byte遍历，如果有汉字，就会有问题
		fmt.Printf("i=%d, val=%c \n", i, strValueRune[i])
	}
	for index, val := range strValue { //按字符遍历
		fmt.Printf("index= %d, val=%c \n", index, val)
	}

	//返回随机数
	rand.Seed(time.Now().UnixNano()) //设置种子
	intn := rand.Intn(100) + 1
	fmt.Println("intn=", intn)
	//break label

label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i= %v, j= %v", i, j)
			break label1
		}
	}

label2:
	for i := 0; i < 10; i++ {
		fmt.Println("continue label i starting")
		for j := 0; j < 10; j++ {
			fmt.Println("continue label j starting")
			continue label2
			fmt.Println("continue label j end")
		}
		fmt.Println("continue label i end")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("goto i=", i)
		goto labelGoto
	}
labelGoto:
	fmt.Println("goto target label")

}
