package main

import "fmt"

// Go 语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：
// 函数 + 引用环境 = 闭包
// 提供一个值, 每次调用函数会指定对值进行累加
func main() {
	// 创建一个累加器, 初始值为 1，accumulator 的类型是 func() int
	accumulator := Accumulate(1)
	// 累加 1 并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为 10
	accumulator2 := Accumulate(10)
	// 累加 1 并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)

	// 创建一个玩家生成器
	generator := playerGen("high noon")
	// 返回玩家的名字和血量
	name, hp := generator()
	// 打印值
	fmt.Println(name, hp)
}

// 返回值类型是 func() int 的函数变量
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value ++
		// 返回一个累加值
		return value
	}
}

// 闭包的记忆效应被用于实现类似于设计模式中工厂模式的生成器，下面的例子创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为 150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}