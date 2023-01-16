package main

//包的引入方式 引入多个的时候的格式
//import(
//"aa"
//"bb"
//)
import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	var i int = 1

	fmt.Println("i=", i)

	// rune 和 byte的使用
	var runeType rune
	runeType = -11121

	var byteType byte
	byteType = 255

	fmt.Println("runeType=", runeType, "byteType=", byteType)
	//Printf按照格式输出
	fmt.Printf("runeType的数据类型 %T, 占用的字节 %d\n", runeType, unsafe.Sizeof(runeType))

	//浮点数  float32 float64
	var floatTest float32 = 1.2
	fmt.Println("floatTest=", floatTest)

	//字符的使用
	var byteTest byte = 'a'
	fmt.Printf("byteTest数据类型 %T, 长度 %d , 值为 %d, 字符值 %c", byteTest, unsafe.Sizeof(byteTest), byteTest, byteTest)
	fmt.Println("")
	//bool 使用
	var boolType bool = true
	fmt.Println("boolType=", boolType, "size=", unsafe.Sizeof(boolType))

	//字符使用:  字符串 %s 字符: %c 数字：%d  浮点数： %f  按照变量的值输出: %v

	var school string
	school = "beijingdaxue"
	fmt.Println("")
	fmt.Printf("school= %s", school)
	fmt.Println("")
	var address string = `\n\t\r`
	fmt.Println(address)

	//基本类型转换  Go中类型转换必须是显示转换 T(v) v转换为T类型
	var iValue int = 100
	var fValue float32 = float32(iValue)
	fmt.Printf("f= %v ", fValue)

	//基本类型转String
	//第一种转换方式 fmt.Sprintf()
	var conValue int64 = 200
	var conStr string = fmt.Sprintf("%d", conValue)
	fmt.Printf("conStr %s \n", conStr)

	//使用 strconv包的函数
	conStrFormat := strconv.FormatInt(conValue, 10)
	fmt.Printf("conStrFormat = %s \n", conStrFormat)

	// string 转 基本类型
	strBase := "12"
	//当函数返回多个值的时候：如果不关心返回的值可以使用 "_" 表示忽略 多个值的情况下，返回的值的接受方式
	strIntValue, _ := strconv.ParseInt(strBase, 10, 64)
	fmt.Printf("strIntValue = %d \n", strIntValue)

	var strFloat = "123.456"
	strFloatValue, _ := strconv.ParseFloat(strFloat, 64)
	fmt.Printf("strFloatValue = %f \n", strFloatValue)

}
