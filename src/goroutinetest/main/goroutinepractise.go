package main

import (
	"fmt"
	"runtime"
	"sync"
)

func PrintHelloWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("PrintHelloWorld hello world :", i)
		//time.Sleep(time.Second)
	}
}

var dataMap sync.Map

func factorial(n int) {

	var value int = 1
	for i := 0; i <= n; i++ {
		value *= i
	}
	dataMap.Store(n, value)
}

func main() {
	// 开启协程
	go PrintHelloWorld()
	// 主线程输出
	for i := 0; i < 10; i++ {
		fmt.Println("main hello world :", i)
		//time.Sleep(time.Second)
	}

	numCpu := runtime.NumCPU() // 获取cpu的数
	fmt.Println("numCpu:", numCpu)
	runtime.GOMAXPROCS(numCpu) // 设置go 使用的最大cpu数

	for i := 1; i < 10; i++ {
		go factorial(i)
	}
	dataMap.Range(func(key, value any) bool {
		fmt.Printf("map key: %v, value: %v \n", key, value)
		return true
	})

}
