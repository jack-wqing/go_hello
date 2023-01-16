package main

import (
	"fmt"
	"sort"
)

func main() {
	// map声明之后必须 make 才能进行 赋值和使用
	var aMap map[string]string         //声明
	aMap = make(map[string]string, 10) //分配内存
	aMap["a"] = "a-value"
	fmt.Println("aMap:", aMap)

	// 使用方式
	//第一种使用方式
	var map1 map[string]string
	map1 = make(map[string]string, 4)
	map1["key1"] = "value1"
	//第二种使用方式
	map2 := make(map[string]string, 4)
	map2["key-map2"] = "value-map2"
	//第三种使用方式
	var map3 = map[string]string{"key1-map3": "value3-map3"}

	fmt.Printf("map init: %v, %v, %v \n", map1, map2, map3)

	delete(map3, "key1-map3")
	fmt.Printf("map init: %v \n", map3)
	//查找
	val, findRes := map1["key1"]
	fmt.Printf("map find: val:%v, findRes:%v \n", val, findRes)
	//map遍历
	for key, value := range map1 {
		fmt.Printf("key:%v, value:%v \n", key, value)
	}

	//map的切片
	var mapSlice []map[string]string = make([]map[string]string, 2)
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string)
		mapSlice[0]["name"] = "猪八戒"
		mapSlice[0]["age"] = "12"
	}
	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string)
		mapSlice[1]["name"] = "孙悟空"
		mapSlice[1]["age"] = "15"
	}
	newEle := make(map[string]string)
	newEle["name"] = "唐僧"
	newEle["age"] = "99"
	mapSlice = append(mapSlice, newEle)
	fmt.Println(mapSlice)
	// map进行排序输出
	var sortMap map[int]int = map[int]int{1: 1, 12: 12, 2: 2, 14: 14, 10: 10}
	var keys []int
	for key, _ := range sortMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%d \t", sortMap[keys[i]])
	}
	fmt.Println()

}
