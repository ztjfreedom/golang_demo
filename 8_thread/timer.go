package main

import (
	"fmt"
	"time"
)

func main() {
	timeAfterFunc()
	tickerTimer()
}

func timeAfterFunc() {
	// 声明一个退出用的通道
	exit := make(chan int)
	// 打印开始
	fmt.Println("start")
	// 过 1 秒后, 调用匿名函数，time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用
	time.AfterFunc(time.Second, func() {
		// 1 秒后, 打印结果
		fmt.Println("one second after")
		// 通知 main() 的 goroutine 已经结束
		exit <- 0
	})
	// 等待结束
	<-exit
}

// 计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发。打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发。这两种方法创建后会返回 time.Ticker 对象和 time.Timer 对象，里面通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），使用这个通道就可以获得时间触发的通知
func tickerTimer() {
	// 创建一个打点器, 每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	// 声明计数变量
	var i int

	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C:  // 计时器到时了
			fmt.Println("stop")
			// 跳出循环
			goto StopHere
		case <-ticker.C:  // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		}
	}

	// 退出的标签, 使用 goto 跳转
StopHere:
	fmt.Println("done")
}