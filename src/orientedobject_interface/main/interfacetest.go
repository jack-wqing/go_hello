package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
)

// Usb 声明一个接口
type Usb interface {
	Start() //声明没有实现的接口
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("手机开始工作。。。")
}

func (p *Phone) Stop() {
	fmt.Println("手机结束工作")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c *Camera) Stop() {
	fmt.Println("相机结束工作")
}

type Computer struct {
}

// Working 实现 Usb 接口(就是指实现来 Usb接口声明的方法)  接口的实例传入只能是指针
func (com *Computer) Working(usb Usb) {
	usb.Start()
}

func (com *Computer) Stopping(usb Usb) {
	usb.Stop()
}

// AInterface 接口直接的继承
type AInterface interface {
	TestO1()
}

type BInterface interface {
	Test02()
}

type CInterface interface {
	AInterface //接口的继承，也是直接匿名引入的方式，和结构体的继承是一样的
	BInterface
	Test03()
}

type T interface {
}

// 接口应用

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (h HeroSlice) Len() int {
	return len(h)
}

func (h HeroSlice) Less(i, j int) bool {
	heroI := h[i]
	heroJ := h[j]
	if heroI.Age > heroJ.Age {
		return false
	} else {
		return true
	}
}

func (h HeroSlice) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}

func main() {

	computer := Computer{}
	phone := &Phone{}
	camera := &Camera{}
	computer.Working(phone)
	computer.Stopping(phone)
	computer.Working(camera)
	computer.Stopping(camera)

	fmt.Printf("OS:%v, ARCH:%v \n", runtime.GOOS, runtime.GOARCH)
	// ------------- 接口应用
	var heroArr HeroSlice = make(HeroSlice, 10)
	for i := 0; i < 10; i++ {
		heroArr[i] = Hero{
			Name: fmt.Sprintf("英雄·%d", i),
			Age:  rand.Intn(100),
		}
	}
	fmt.Println("排序前：", heroArr)
	sort.Sort(heroArr)
	fmt.Println("排序后：", heroArr)

}
