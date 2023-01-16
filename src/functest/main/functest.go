package main

import (
	"fmt"
	packUtil "go_hello/src/functest/pactest" //自定义的包，从项目路径引入,可以给包起别名
)

func init() {
	fmt.Println("main init test ...")
}

func main() {

	var num1 int
	var num2 int
	var operator byte
	for {
		fmt.Println("请输入第一个数")
		fmt.Scanln(&num1)
		fmt.Println("请输入第二个数")
		fmt.Scanln(&num2)
		fmt.Println("请输入操作符")
		fmt.Scanln(&operator)
		result := packUtil.Calculate(num1, num2, operator)
		fmt.Printf("result=%d \n", result)
	}

}
