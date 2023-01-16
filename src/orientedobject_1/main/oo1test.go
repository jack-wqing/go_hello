package main

import (
	"encoding/json"
	"fmt"
)

// Cat 定义一个结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

type Point struct {
	x int
	y int
}

// Rect 结构体属性内存分布
type Rect struct {
	leftTop, rightDown Point
}

type A struct {
	a int
}

type B struct {
	a int
}

type Monster struct {
	Name  string `json:"name"` //结构体的标签
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	// 创建结构体实例，默认具有值来
	var catObj Cat
	// 赋值方式通过 "." 操作符
	catObj.Name = "小白"
	catObj.Age = 12
	catObj.Color = "白色"
	fmt.Println("catObj init:", catObj)
	fmt.Println("cat Color:", catObj.Color)

	var catObj1 Cat = Cat{} //创建方式
	fmt.Println("catObj1:", catObj1)

	var a A = A{}
	var b B = B(a)
	fmt.Printf("B convert: %v \n", b)

	var monster Monster = Monster{"牛魔王", 500, "芭蕉扇"}
	marshal, _ := json.Marshal(monster) // struct to json string  使用的反射功能
	fmt.Printf("monster json:%s", marshal)

}
