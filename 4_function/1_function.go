package main

import (
	"fmt"
	"time"
)

func main() {
	result := add(2, 1)
	fmt.Println(result)

	// 函数变量
	subFunc := sub
	fmt.Println(subFunc(2, 1))

	fmt.Println(twoReturnValues())

	// 匿名函数，(100) 表示对匿名函数进行调用，传递参数为 100
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(100)

	// 匿名函数作为 callback
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()

	optionalParam(1, 2, 3)

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	optionalParamType(v1, v2, v3, v4)

	fmt.Println(fibonacci(10))

	runningTime()
}

/*
  func 函数名(形式参数列表) (返回值列表) {
	  函数体
  }
 */
func add(x int, y int) int {
	return x + y
}

// 带有变量名的返回值
func sub(x int, y int) (z int) {
	z = x - y
	return
}

func twoReturnValues() (int, string) {
	return 1, "hello"
}

// 遍历切片的每个元素, 通过指定处理函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

// 可变参数
func optionalParam(args ...int) {
	for _, arg := range args {
		fmt.Printf("%d ", arg)
	}
	fmt.Println()
}

// 任意类型的可变参数，用 interface{} 传递任意类型数据是 Go 语言的惯例用法，使用 interface{} 仍然是类型安全的
func optionalParamType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
	// 可变参数的传递
	optionalParamTypePrint(args...)
}

func optionalParamTypePrint(args ...interface{}) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}

// 递归 + dp
var fibs [100]int64
func fibonacci(n int) (res int64) {
	// 记忆化：检查数组中是否已知斐波那契（n）
	if fibs[n] != 0 {
		return fibs[n]
	}
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func runningTime() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	// 在 Go 语言中我们可以使用 time 包中的 Since() 函数来获取函数的运行时间
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}