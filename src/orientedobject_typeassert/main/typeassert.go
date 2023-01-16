package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {

	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point

	b = a.(Point) // 表示类型断言， 类型之间进行转换
	fmt.Println("b=", b)
	//带检测的类型断言

	var i, ok = a.(int)
	if ok { // 表示断言成功， 转换成功
		fmt.Printf("convert success: i value : %v \n", i)
		fmt.Printf("a type: %T \n", a)
	} else {
		fmt.Println("转换失败")
		fmt.Printf("a type: %T not convert int \n", a)
	}
	switch a.(type) { //a.(type) 只能使用在switch

	}
}
