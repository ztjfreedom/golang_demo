package main

import "fmt"

func main() {
	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.Add(2))
}

// 将 int 定义为 MyInt 类型
type MyInt int
// type MyInt = int 是类型别名

// 为 MyInt 添加 IsZero() 方法
func (m MyInt) IsZero() bool {
	return m == 0
}
/*
  也可以用指针类型的接收器
  func (m *MyInt) IsZero() bool {
      return *m == 0
  }
 */

// 为 MyInt 添加 Add() 方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}