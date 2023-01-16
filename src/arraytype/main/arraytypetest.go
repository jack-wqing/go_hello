package main

import "fmt"

func main() {

	var array [6]float64 // 数组的定义
	array[0] = 12        // 数组的赋值
	array[1] = 14.2
	total := 0.0
	for i := 0; i < len(array); i++ { //遍历数组
		total += array[i]
	}
	avgValue := total / float64(len(array))
	fmt.Printf("total: %v, avgValue: %.2f \n", total, avgValue)
	//数组的地址：
	fmt.Printf("array position: %p, 第一个元素的地址：%p \n", &array, &array[0])

	//数组的定义方方式
	var arr [3]int = [3]int{1, 2, 3}
	var arr1 = [3]int{4, 5, 6}
	var arr2 = [...]int{7, 8, 9}
	var arrStr = [3]string{1: "tom", 2: "jack", 0: "jerry"}
	fmt.Printf("array init: arr:%v, arr1:%v, arr2:%v, arrStr:%v \n", arr, arr1, arr2, arrStr)
	//通过for range遍历 array
	for index, value := range arrStr {
		fmt.Printf("for range: index:%v, value:%v \n", index, value)
	}
}
