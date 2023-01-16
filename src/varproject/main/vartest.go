package main

import "fmt"

// 第一种：定义全局变量
var globalAge = 12
var globalName = "jack"

// 第二种：定义全局变量
var (
	globalSchool  = "lanzhoudaxue"
	globalFriends = 12
)

func main() {
	//定义变量
	var i int
	//赋值变量
	i = 10
	//使用变量
	fmt.Println("i=", i)
	//类型自动推导
	var b = 10
	fmt.Println("b=", b)
	//省略var, 注意 := 左侧的变量不应该是已经声明过的变量，否则回导致编译错误
	tom := "tom"
	fmt.Println("name=", tom)

	//多变量声明
	var n1, n2, n3 int
	n3 = 12
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)

	//多变量声明不同的类型
	var m1, name, m2 = 100, "tom", 888
	fmt.Println("m1=", m1, "name=", name, "m2=", m2)

	//多变量类型推导 值与声明之间必须要有一个赋值的分割
	q1, q2, q3 := "12", "23", "34"
	fmt.Println("q1=", q1, "q2=", q2, "q3=", q3)
	//全局变量使用
	fmt.Println("global_age=", globalAge, "global_name=", globalName, "global_school=", globalSchool, "global_friends=", globalFriends)

	var num1, num2 = 2, 1
	num3 := num1 + num2
	var s1, s2 = "hello", "jack"
	str3 := s1 + s2
	fmt.Println("num3=", num3, "str3=", str3)

}
