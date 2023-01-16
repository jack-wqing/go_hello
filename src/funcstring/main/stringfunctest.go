package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串函数详解
func main() {

	//字符串长度 len是内建函数
	lenstr := "abcde"
	fmt.Printf("len value : %d \n", len(lenstr))
	//遍历字符，如果包含中文， 切片的方式 rune的切片
	strRuneArr := []rune(lenstr)
	for i := 0; i < len(strRuneArr); i++ {
		if i == 0 {
			fmt.Printf("[]rune str value: %c", strRuneArr[i])
		} else {
			fmt.Printf("\t%c", strRuneArr[i])
		}
		if i == len(strRuneArr) {
			fmt.Printf("\n")
		}
	}

	//字符转int
	var str = "123"
	istr, _ := strconv.Atoi(str)
	fmt.Printf("string to int: %d \n", istr)
	// int 转 string
	var intValue = 123
	var intStr = strconv.Itoa(intValue)
	fmt.Printf("int to str : %s \n", intStr)

	//str 转 byte
	var byteStr = "abcd中国"
	bytesArray := []byte(byteStr)
	for i := 0; i < len(bytesArray); i++ {
		fmt.Printf("%c \t", bytesArray[i])
	}
	fmt.Println("")
	//byte to str
	var byteValue = []byte{97, 98, 99}
	byteStr = string(byteValue)
	fmt.Printf("byte to str %v \n", byteStr)
	//进制转换
	intFormatBinary := strconv.FormatInt(12, 2)
	fmt.Printf("int format binary: %T %v \n", intFormatBinary, intFormatBinary)
	// 字符串strings包的使用
	strContain := strings.Contains("abca", "a")
	fmt.Printf("strings Contains: %v \n", strContain)
	stringsCount := strings.Count("aaa", "a")
	fmt.Printf("strings count: %v \n", stringsCount)

	var stringsEqualFold bool = strings.EqualFold("abc", "ABC")
	fmt.Printf("strings EqualFold: %v \n", stringsEqualFold)
	var stringsIndex = strings.Index("abc", "c")
	fmt.Printf("strings Index: %v \n", stringsIndex)
	//判断字符串开头和结尾
	var hasPrefixValue = strings.HasPrefix("abcd", "ab")
	fmt.Printf("hasPrefixValue: %v \n", hasPrefixValue)

	var hasSuffixValue = strings.HasSuffix("abcde", "cde")
	fmt.Printf("hasSuffixValue: %v \n", hasSuffixValue)

}
