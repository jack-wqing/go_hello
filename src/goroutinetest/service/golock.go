package service

import (
	"fmt"
	"sync"
	"time"
)

var (
	dataMap map[int]uint64 = make(map[int]uint64)
	lock    sync.Mutex     //声明一个全局的互斥锁， lock是一个全局的互斥锁
)

func Factorial(n int) {

	var value uint64 = 1
	for i := 1; i < n; i++ {
		value *= uint64(i)
	}
	lock.Lock() // 加锁
	dataMap[n] = value
	lock.Unlock() // 解锁

}

func PrintDataMap() {

	time.Sleep(1 * time.Second)
	lock.Lock()
	for key, value := range dataMap {
		fmt.Printf("test key: %v, value: %v \n", key, value)
	}
	lock.Unlock()

}
