package main

import "fmt"

// nil 在 Go语言中只能被赋值给指针和接口。接口在底层的实现有两个部分：type 和 data。在源码中，显式地将 nil 赋值给接口时，接口的 type 和 data 都将为 nil。此时，接口与 nil 值判断是相等的
// 但如果将一个带有类型的 nil 赋值给接口时，只有 data 为 nil，此时，接口与 nil 判断将不相等
func main() {
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}

	if GetStringerNil() == nil {
		fmt.Println("GetStringerNil() == nil")
	} else {
		fmt.Println("GetStringerNil() != nil")
	}
}

// 定义一个结构体
type MyImplement struct{}

// 实现fmt.Stringer的String方法
func (m *MyImplement) String() string {
	return "hi"
}

// 在函数中返回 fmt.Stringer 接口
func GetStringer() fmt.Stringer {
	// 赋nil
	var s *MyImplement = nil
	// 返回变量
	return s
}

// 为了避免这类误判的问题，可以在函数返回时，发现带有 nil 的指针时直接返回 nil
func GetStringerNil() fmt.Stringer {
	var s *MyImplement = nil
	if s == nil {
		return nil
	}
	return s
}