package model

type Student struct {
	Name  string
	Score float64
}

type teach struct {
	Name string
	Age  int
}

// NewTeach 工厂模式来结果 结构体变量小写
func NewTeach(name string, age int) *teach {
	return &teach{name, age}
}
