package main

import (
	"errors"
	"fmt"
	"time"
)

const ( //定义常量
	Nanosecond  = 1                  //纳秒
	Microsecond = 1000 * Nanosecond  //微妙
	Millisecond = 1000 * Microsecond //毫秒
	Second      = 1000 * Millisecond //秒
	Minute      = 60 * Second        //分钟
	Hour        = 60 * Minute        //小时
)

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 1
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "conf.ini" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func test01(name string) {
	err := readConf(name)
	if err != nil {
		panic(err)
	}
	fmt.Println("test01 自定义代码继续执行 name:", name)
}

func main() {

	//日期和时间
	nowTime := time.Now()
	fmt.Printf("time type:%T \t %v \n", nowTime, nowTime)
	//时间的各个组成部分
	fmt.Println("Year:", nowTime.Year())
	fmt.Println("Month:", nowTime.Year())
	fmt.Println("Day:", nowTime.Year())
	fmt.Println("Hour:", nowTime.Hour())
	fmt.Println("Minute:", nowTime.Minute())
	fmt.Println("Second:", nowTime.Second())

	fmt.Printf("当前年月日： %d-%d-%d %d:%d:%d", nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Month(), nowTime.Second())
	ymdhmsFormat := fmt.Sprintf("当前年月日： %d-%d-%d %d:%d:%d", nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Month(), nowTime.Second())
	fmt.Printf("ymdhms:%v \n", ymdhmsFormat)
	timeFormat := nowTime.Format("2006/01/02 15+04+05")
	fmt.Printf("timeFormat: %v \n", timeFormat)

	//休眠时间
	fmt.Println("start time:", time.Now().Format("2006-01-02 15:04:05"))
	//time.Sleep(2 * time.Second) //当前程序睡眠 2秒
	fmt.Println("start time:", time.Now().Format("2006-01-02 15:04:05"))
	//秒数： time.Unix()  纳秒：time.unixNano()
	fmt.Println("time.Unix:", time.Now().Unix())
	fmt.Println("time.UnixNano:", time.Now().UnixNano())
	// new 内置函数
	newInt := new(int)
	fmt.Printf("newInt type: %T, address:%v, value: %v \n", newInt, newInt, *newInt)
	fmt.Println("异常处理")
	test()
	fmt.Println("异常之后的代码")

	//自定义错误
	test01("conf.ini")
	test01("config")

}
