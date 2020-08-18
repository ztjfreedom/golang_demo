package main

import (
	"fmt"
	"golearn/v2/7_package/model"
)

func main() {
	// 根据字符串动态创建一个 Class1 实例
	c1 := model.Create("Class1")
	c1.Do()
	// 根据字符串动态创建一个 Class 2实例
	c2 := model.Create("Class2")
	c2.Do()
}

// 定义类 1
type Class1 struct {
}

// 实现 Class 接口
func (c *Class1) Do() {
	fmt.Println("Class1")
}

// 定义类 2
type Class2 struct {
}

// 实现 Class 接口
func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {
	// 在启动时注册类 1 工厂
	model.Register("Class1", func() model.Class {
		return new(Class1)
	})

	// 在启动时注册类 2 工厂
	model.Register("Class2", func() model.Class {
		return new(Class2)
	})
}