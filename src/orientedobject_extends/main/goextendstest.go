package main

import "fmt"

type Goods struct {
	Name  string
	Price int
}

func (g Goods) showInfo() {
	fmt.Printf("Goods Name: %v, Price: %v", g.Name, g.Price)
}

type Book struct {
	Goods
	Writer string
}

type Diary struct {
	good    Goods
	Address string
}

type BiJi struct {
	*Goods //指针的形式
	int    //匿名基本类型   用的时候直接 BiJi.int
}

func main() {

	var book Book = Book{Goods{Name: "西游记", Price: 12}, "吴承恩"}
	fmt.Printf("book:%v \n", book)
	fmt.Printf("book name:%v \n", book.Name)
	fmt.Printf("book price:%v \n", book.Price)
	fmt.Printf("book writer:%v \n", book.Writer)
	book.Name = "三国"
	book.Price = 12
	book.Writer = "..."
	fmt.Printf("book:%v \n", book)
	fmt.Printf("book name:%v \n", book.Name)
	fmt.Printf("book price:%v \n", book.Price)
	fmt.Printf("book writer:%v \n", book.Writer)
	book.showInfo()

	diary := Diary{good: Goods{"haha", 12}, Address: "www"}
	diary.good.Name = "good" //组合结构体

}
