package main

import "fmt"

func main() {
	receiveFromClosedChan()
	sendToClosedChan()
}

// 被关闭的通道不会被置为 nil。如果尝试对已经关闭的通道进行发送，将会触发宕机
func sendToClosedChan() {
	// 创建一个整型的通道
	ch := make(chan int)
	// 关闭通道
	close(ch)
	// 打印通道的指针, 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))
	// 给关闭的通道发送数据
	ch <- 1
}

/*
  从已关闭的通道接收数据时将不会发生阻塞
  运行结果前两行正确输出带缓冲通道的数据，表明缓冲通道在关闭后依然可以访问内部的数据
  运行结果第三行的“0 false”表示通道在关闭状态下取出的值。0 表示这个通道的默认值，false 表示没有获取成功
 */
func receiveFromClosedChan() {
	// 创建一个整型带两个缓冲的通道
	ch := make(chan int, 2)

	// 给通道放入两个数据
	ch <- 0
	ch <- 1

	// 关闭缓冲
	close(ch)
	// 遍历缓冲所有数据, 且多遍历1个
	for i := 0; i < cap(ch) + 1; i++ {
		// 从通道中取出数据
		v, ok := <-ch
		// 打印取出数据的状态
		fmt.Println(v, ok)
	}
}