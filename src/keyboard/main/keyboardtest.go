package main

import "fmt"

func main() {
	//使用 fmt.Scanln 获取键盘输入
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名 ")
	fmt.Scanln(&name) //当程序执行到这里，回停止等待输入

	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水 ")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过 ")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%v, 年龄: %v, 薪水：%v, 是否通过：%v \n", name, age, sal, isPass)

	//方式2，按指定的格式输入：
	fmt.Println("请输入姓名，年龄，薪水，是否通过 用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%v, 年龄: %v, 薪水：%v, 是否通过：%v \n", name, age, sal, isPass)

	var i int = 5
	var j int = 011
	var k int = 0x11
	fmt.Printf("i=%v j=%v k=%v", i, j, k)
}
