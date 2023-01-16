package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 循环接受客户端发送的数据
	defer conn.Close()
	for {
		var buf []byte = make([]byte, 1024)
		fmt.Println("server waiting client write, client address:", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //会阻塞
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit")
			}
			fmt.Println("client conn.Read err=", err)
			return
		}
		fmt.Printf("client read data length: %v \n", n)
		fmt.Printf("server recive data: %s \n", string(buf[:n]))
	}
}

func main() {
	fmt.Println("开启网络监听")
	// 1、开启接口监听
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	fmt.Printf("listen: %v \n", listen)
	defer listen.Close()
	for {
		//2、等待链接
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("Accept() err1=", err1)
			continue
		}
		fmt.Printf("Accept() sun con=%v, client addr:%v \n", conn, conn.RemoteAddr().String())
		// 3、启动协程处理客户端
		go process(conn)
	}

}
