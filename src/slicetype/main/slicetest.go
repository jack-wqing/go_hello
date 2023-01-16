package main

import (
	"fmt"
)

// 二分查找
func binarySearchFun(data *[6]int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	var middle = (end + start) / 2
	if target == (*data)[middle] {
		return middle
	} else if target < data[middle] {
		return binarySearchFun(data, 0, middle-1, target)
	} else {
		return binarySearchFun(data, middle+1, end, target)
	}

}

func main() {

	// 切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 表示下标 1 到 3，但是不包含下标为3的元素
	var slice []int = intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素是 =", slice)
	fmt.Println("slice length =", len(slice)) // 元素长度
	fmt.Println("slice cap=", cap(slice))     // 容量

	fmt.Printf("slice type : %T \n", slice)
	fmt.Printf("intArr[1] address: %v, slice[0]:%v \n", &intArr[1], &slice[0])

	//通过 make 创建切片
	var makeSlice []int = make([]int, 4, 10)
	fmt.Println("makeSlice:", makeSlice)
	fmt.Printf("slice len=%v, cap=%v \n", len(makeSlice), cap(makeSlice))
	makeSlice[0] = 11
	makeSlice[1] = 12
	fmt.Println("makeSlice update:", makeSlice)

	//创建切片直接赋值
	var sliceInit []int = []int{1, 2, 3}
	fmt.Printf("sliceInit: len: %v, cap: %v, value:%v \n", len(sliceInit), cap(sliceInit), sliceInit)
	for i := 0; i < len(sliceInit); i++ {
		fmt.Printf("%d \t", sliceInit[i])
	}
	fmt.Println("")
	for index, value := range sliceInit {
		fmt.Printf("%d,%d \t", index, value)
	}
	fmt.Println("")

	var sliceArr = [...]int{1, 2, 3, 4}
	var sliceFirst = sliceArr[:3]
	var sliceSecond = sliceArr[2:]
	var sliceThird = sliceArr[:]
	fmt.Printf("slice Array:sliceFirst:%d, sliceSecond:%d, sliceThird:%d \n", sliceFirst, sliceSecond, sliceThird)
	//切片的动态变化
	var sliceAppend = []int{1, 2, 3}
	sliceAppend = append(sliceAppend, 4) //返回的是新的slice
	fmt.Printf("sliceAppend:%v \n", sliceAppend)
	var sliceOrigin = []int{4}
	sliceOrigin = append(sliceOrigin, sliceAppend...) //三个点表示取出元素
	fmt.Printf("sliceOrigin:%v \n", sliceOrigin)

	var sliceCopy = make([]int, 10)
	copy(sliceCopy, sliceOrigin)
	fmt.Println("sliceCopy:", sliceCopy)

	//string 和 slice
	var sliceStr string = "hello world 好的"
	var sliceStrArr = sliceStr[:3] // 这样是字符串
	fmt.Printf("sliceStrArr:%T, value:%v \n", sliceStrArr, sliceStrArr)
	var sliceStrByte []byte = []byte(sliceStr)
	fmt.Printf("sliceStrByte:%T, value:%v \n", sliceStrByte, sliceStrByte)
	// 字符串的修嘎
	sliceStrByte[0] = 'Z'
	fmt.Printf("sliceStrByte:%s \n", sliceStrByte)
	//可以将 字符串转为rune来处理中文 按字符操作的
	var sliceRune = []rune(sliceStr)
	sliceRune[0] = '呀'
	fmt.Printf("sliceRune:%c \n", sliceRune)
	//冒泡排序
	var bubbleSortArr = [...]int{1, 2, 3, 4, 9, 4, 8, 5}
	for i := 0; i < len(bubbleSortArr); i++ {
		for j := 1; j < len(bubbleSortArr)-i; j++ {
			if bubbleSortArr[j-1] > bubbleSortArr[j] {
				temp := bubbleSortArr[j-1]
				bubbleSortArr[j-1] = bubbleSortArr[j]
				bubbleSortArr[j] = temp
			}
		}
	}
	fmt.Println("bubbleSortArr:", bubbleSortArr)
	//顺序查找，二分查找
	var binarySearch = [6]int{1, 34, 46, 23, 11, 78}

	binarySearchIndex := binarySearchFun(&binarySearch, 0, len(binarySearch), 46)
	fmt.Println("binarySearchIndex:", binarySearchIndex)

	//二维数组
	var arrArr [4][6]int
	arrArr[1][2] = 1
	arrArr[2][1] = 2
	arrArr[2][3] = 3
	fmt.Printf("arrArr:%v /n", arrArr)
	//二维数组遍历
	for _, value := range arrArr {
		for _, ele := range value {
			fmt.Printf("%v ", ele)
		}
		fmt.Println("")
	}

}
