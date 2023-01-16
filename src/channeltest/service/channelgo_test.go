package service

import (
	"fmt"
	"testing"
)

func TestChannelReadWrite(t *testing.T) {

	go WriteData()
	go ReadData()
	channel := GetExitChannel()
	for {
		v, ok := <-channel
		if ok {
			fmt.Printf("test main: %v, status: %v \n", v, ok)
			break
		}
	}
	fmt.Println("pass")
}

func TestSelectTest(t *testing.T) {
	SelectTest()
}

func TestAllTest(t *testing.T) {
	AllTest()
	for i := 0; i < 20; i++ {
		fmt.Printf("testing hello world : %v \n", i)
	}
}
