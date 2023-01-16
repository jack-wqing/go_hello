package service

import "testing"

func TestChannelTest(t *testing.T) {
	ChannelTest() // 测试关闭之后，向管道中写东西
}

func TestChannelTest1(t *testing.T) {
	ChannelTest1()
}

func TestChannelOnlyRead(t *testing.T) {
	ChannelOnlyRead()
}
