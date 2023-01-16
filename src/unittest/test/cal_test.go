package test

import (
	"go_hello/src/unittest/service"
	"testing"
)

func TestAddUpper(t *testing.T) {

	addUpper := service.AddUpper(10)

	if addUpper == 55 {
		t.Logf("AddUpper(10)计算正确：期望值：%v, 实际值: %v \n", 55, addUpper)
	} else {
		t.Fatalf("AddUpper(10)计算错误：期望值：%v, 实际值: %v \n", 55, addUpper)
	}

}

func TestSub(t *testing.T) {
	sub := service.Sub(1, 2)
	t.Logf("sub(1, 2) value: %v", sub)
}
