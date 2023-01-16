package main

import (
	"fmt"
)

func main() {

	var i int = 100

	fmt.Println("i的地址：", &i)
	//定义指针
	var ptr *int = &i
	//获取指针变量的值
	fmt.Sprintf("ptr的地址 %d \n", &ptr)
	fmt.Printf("ptr value: %v \n", ptr)
	fmt.Println("ptr* = ", *ptr)

}
