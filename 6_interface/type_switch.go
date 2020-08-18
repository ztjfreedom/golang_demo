package main

import "fmt"

func main() {
	printType(1024)
	printType("pig")
	printType(true)
	printInterface(new(Alipay))
	printInterface(new(Cash))
}

// type-switch 流程控制的语法或许是Go语言中最古怪的语法。 它可以被看作是类型断言的增强版。它和 switch-case 流程控制代码块有些相似
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

// 使用类型分支判断接口类型
// 打印支付方式具备的特点
func printInterface(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainCanUseFaceID:  // 可以刷脸
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:  // 可能被偷
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

// 电子支付方式
type Alipay struct {
}

// 为 Alipay 添加 CanUseFaceID() 方法, 表示电子支付方式支持刷脸
func (a *Alipay) CanUseFaceID() {
}

// 现金支付方式
type Cash struct {
}

// 为 Cash 添加 Stolen() 方法, 表示现金支付方式会出现偷窃情况
func (a *Cash) Stolen() {
}

// 具备刷脸特性的接口
type ContainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}