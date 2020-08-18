package main

import (
	"errors"
	"fmt"
	"time"
)

/*
  Go 语言中提供了 select 关键字，可以同时响应多个通道的操作。select 的用法与 switch 语句非常类似，由 select 开始一个新的选择块，每个选择条件由 case 语句来描述
  与 switch 语句可以选择任何可使用相等比较的条件相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作，大致结构如下：
  select{
      case 操作1:
          响应操作1
      case 操作2:
          响应操作2
      …
      default:
          没有操作情况
  }
 */
func main() {
	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	// 并发执行服务器逻辑
	go RPCServer(ch)

	// 客户端请求数据和接收数据
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		// 发生错误打印
		fmt.Println(err)
	} else {
		// 正常接收到数据
		fmt.Println("client received:", recv)
	}
}

// 服务器开发中会使用 RPC（Remote Procedure Call，远程过程调用）简化进程间通信的过程。RPC 能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数调用一样
// 模拟 RPC 客户端的请求和接收消息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case ack := <-ch: // 接收到服务器返回数据
		return ack, nil
	case <-time.After(time.Second): // 超时
		return "", errors.New("Time out")
	}
}

// 模拟 RPC 服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch

		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 反馈给客户端收到
		ch <- "roger"
	}
}