package service

import "fmt"

func ChannelTest() {

	var intChan chan int = make(chan int, 5)
	intChan <- 10
	intChan <- 23
	close(intChan) // 通过内置的 close 关闭 channel
	intChan <- 33
	for value := range intChan {
		fmt.Printf("channel value: %v \n", value)
	}

}

func ChannelTest1() {

	var intChan chan int = make(chan int, 5)
	intChan <- 10
	intChan <- 23
	intChan <- 33
	close(intChan)
	for value := range intChan { // for range 只能遍历关闭之后的 channel
		fmt.Printf("channel value: %v \n", value)
	}

}

func ChannelOnlyRead() {
	var rwChannel chan int //可读可写 channel
	rwChannel = make(chan int, 1)
	rwChannel <- 1
	num := <-rwChannel
	close(rwChannel)
	fmt.Printf("rwChannel num type: %T, value: %v \n", num, num)
	var wChannel chan<- int //只写 channel
	wChannel = make(chan<- int, 3)
	wChannel <- 23
	close(wChannel)
	//var rChannel <-chan int //只读 channel
	//rChannel = make(<-chan int)
	//num1 := <-rChannel
	//fmt.Println("rChannel:", num1)

}
