package main

import (
	"fmt"
)

func main() {
	// 创建一个channel
	c := make(chan int)

	// 并发执行 printer, 传入 channel
	go printer(c)

	for i := 1; i <= 10; i++ {
		// 将数据通过 channel 投送给 printer
		c <- i
	}

	// 通知并发的 printer 结束循环
	c <- 0

	// 等待 printer 结束
	<-c
}

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从 channel 中获取一个数据
		data := <-c

		// 将 0 视为数据结束
		if data == 0 {
			break
		}

		// 打印数据
		fmt.Println(data)
	}

	// 通知main已经结束循环
	c <- 0
}