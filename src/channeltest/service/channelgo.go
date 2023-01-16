package service

import (
	"fmt"
	"time"
)

var intChan chan int = make(chan int, 5)

var exitChan chan bool = make(chan bool, 1)

func WriteData() {
	for i := 0; i < 20; i++ {
		intChan <- i
		time.Sleep(time.Second)
		fmt.Printf("channel writeData value: %v \n", i)
	}
	close(intChan)
}

func ReadData() {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("channel readData value: %v \n", v)
	}
	exitChan <- true
	close(exitChan)
}

func GetExitChannel() chan bool {
	return exitChan
}

func SelectTest() {
	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf(" %d", i)
	}
	// 传统的传统的channel，不关闭，编程 deadlock问题
	// 问题：实际不好确定什么时候关闭管道
	// 使用 select 方式解决
	for {
		select {
		case iv := <-intChan:
			fmt.Println("intChan value: ", iv)
			time.Sleep(100 * time.Millisecond)
		case sv := <-strChan:
			fmt.Println("strChan value: ", sv)
			time.Sleep(100 * time.Millisecond)
		default:
			fmt.Println("都取不到数据")
			return
		}
	}
}

func PanicGoTest() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func MapInitTest() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("map init test error: ", err)
		}
	}()

	var mapData map[int]string
	mapData[1] = "hello world"
}

func AllTest() {
	go MapInitTest()
	go PanicGoTest()
}
