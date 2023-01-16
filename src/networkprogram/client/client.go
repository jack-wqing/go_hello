package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1、链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
	}
	fmt.Printf("net.Dial suc conn=%v \n", conn)
	// 2、客户端输入
	reader := bufio.NewReader(os.Stdin) // os.Stdin 标准输入  os.Stdout  os.Stderr
	for {
		//读取终端的输入
		line, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("readString err1=", err1)
		}
		line = strings.Trim(line, "\n")
		if line == "exit" {
			fmt.Println("client input exit")
			break
		}
		n, err2 := conn.Write([]byte(line))
		if err2 != nil {
			fmt.Println("conn.Write err2=", err2)
		}
		fmt.Println("write number byte n=", n)
	}
	conn.Close()
}
