package main

import (
	"fmt"
	"os"
	"sync"
)

/*
  Go 语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行
  关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件
*/
func main()  {
	deferOrder()
	fmt.Println(readValue("key"))
	fmt.Println(fileSize("/Users/zoutaojie/Workspace/GoLearn/3_control/article.txt"))
}

func deferOrder() {
	fmt.Println("defer begin")
	// 将 defer 放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}

/*
  处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源
  defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题
 */
var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)


func readValue(key string) int {
	valueByKey["key"] = 100
	valueByKeyGuard.Lock()
	// defer 后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// 延迟调用 Close, 此时 Close 不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		// defer 机制触发, 调用 Close 关闭文件
		return 0
	}
	size := info.Size()
	// defer 机制触发, 调用 Close 关闭文件
	return size
}