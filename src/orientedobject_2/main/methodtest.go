package main

import (
	"fmt"
	"go_hello/src/orientedobject_2/model"
)

type A struct {
	Num int
}

func (a A) test() {
	fmt.Println("A.num=", a.Num)
}

func main() {

	var a A
	a.test()

	var a1 A = A{Num: 12} // 赋值的时候可以指定参数名
	fmt.Println("a1=", a1)

	var student model.Student = model.Student{Name: "tom", Score: 12}
	fmt.Printf("student Name: %v, score: %v \n", student.Name, student.Score)

	var teach = model.NewTeach("jack", 12)

	fmt.Printf("tearch type: %T, teach:%v \n", teach, teach.Name)

	accountOne := model.NewAccount("no1", "no1")
	accountOne.SetBalance(123)
	fmt.Println("accountONe: account.balance:", accountOne.GetBalance())

}
