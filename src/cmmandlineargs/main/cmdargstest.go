package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for index, value := range args {
		fmt.Printf("第 %v 个参数，值为 %v \n", index, value)
	}

	var user string
	var pwd string
	var host string
	var port int
	// &表示接受用户 -u 后面输入的参数值
	// "u", 就是表示 -u 指定的参数
	// ""， 默认值
	// "usage": 值说明
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "pwd", "", "用户名默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名,默认是localhost")
	flag.IntVar(&port, "port", 3306, "端口号默认为3306")
	//必须在所有flag都注册好而未访问其值时执行
	flag.Parse()
	fmt.Printf("用户名为：%v, 密码为:%v, 主机为:%v, 端口为:%v \n", user, pwd, host, port)

}
