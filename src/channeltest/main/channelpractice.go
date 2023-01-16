package main

import "fmt"

func main() {
	//声明管道
	var intChan chan int = make(chan int, 3)
	fmt.Printf("intChan type: %v, 自己的地址：%p \n", intChan, &intChan)

	// 管道中添加数据不能超过容量
	intChan <- 12
	num := 211
	intChan <- num
	fmt.Printf("channel len: %d, cap:%d \n", len(intChan), cap(intChan))
	intChan <- 24
	fmt.Printf("channel len: %d, cap:%d \n", len(intChan), cap(intChan))
	num1 := <-intChan
	fmt.Printf("channel num1: %d \n", num1)
	fmt.Printf("channel len: %d, cap:%d \n", len(intChan), cap(intChan))
	// 没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock错误

}
