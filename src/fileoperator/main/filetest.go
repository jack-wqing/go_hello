package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("src/fileoperator/main/1.txt")
	if err != nil {
		fmt.Println("file open err:", err)
	}
	fmt.Println("file:", file)
	fmt.Println("file name:", file.Name())
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件错误")
	} else {
		fmt.Println("文件已经正常关闭")
	}
	//读文件
	file, err = os.Open("src/fileoperator/main/1.txt")
	//创建Reader
	reader := bufio.NewReader(file)

	for {
		readString, _ := reader.ReadString('\n')
		if readString == "" {
			break
		}
		fmt.Print(readString)
	}
	fmt.Println()
	file.Close() //关闭资源，负责内存泄漏

	dataByte, err1 := os.ReadFile("src/fileoperator/main/1.txt")
	if err1 == nil {
		fmt.Printf("file byte data: %s \n", dataByte)
	} else {
		fmt.Println("os.ReadFile read file error")
	}

	writerFile, err2 := os.OpenFile("src/fileoperator/main/2.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err2 != nil {
		fmt.Println("os.OpenFile error:", err)
		return
	}
	writer := bufio.NewWriter(writerFile)
	writer.WriteString("hello new Writer")
	writer.Flush() // 缓存中的内容需要刷盘
	writerFile.Close()

	stat, err3 := os.Stat("src/fileoperator/main/2.txt")
	if err3 != nil {
		if os.IsNotExist(err3) {
			fmt.Println("文件目录不存在")
		} else {
			fmt.Println("其他错误：err:", err3)
		}
	} else {
		fmt.Printf("FileInfo:%v", stat)
	}

}
