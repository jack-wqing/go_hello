package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b) // reflect.Type
	fmt.Println("rType: ", rType)
	fmt.Printf("rType type: %T \n", rType)
	rVal := reflect.ValueOf(b) // reflect.Value
	fmt.Println("rVal:", rVal)
	fmt.Printf("rVal Value: %T, \n", rVal)

	rI := rVal.Interface() // interface{}
	num2 := rI.(int)
	fmt.Println("num2:", num2)
}

type Student struct {
	Name string
	Age  int
}

// 测试结构体的反射
func reflectStruct(b interface{}) {
	fmt.Println("-----------------------------------------------")
	rType := reflect.TypeOf(b)
	fmt.Printf("rType value: %v, rtype real type: %T \n", rType, rType)
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue value: %v, rType real value: %T \n", rValue, rValue)

	kind := rType.Kind()
	fmt.Printf("kind value: %v, kind Type: %T \n", kind, kind)

	// reflect.value -> interface{}
	iv := rValue.Interface()
	fmt.Printf("iv type: %T, iv value: %v \n", iv, iv)
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu type: %T, stu value: %v \n", stu, stu)
	} else {
		fmt.Println("iv 不是 student 类型")
	}
	const name = iota

}

func main() {
	// 对(基本类型、interface()、reflect.Value, reflect.Type)进行反射的基本操作
	var num int = 20
	reflectTest01(num)
	// 结构体的操作
	var stu = Student{
		Name: "tom",
		Age:  12,
	}
	reflectStruct(stu)

}
