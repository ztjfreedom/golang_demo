package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
  服务端程序中包含 4 个 goroutine，分别是一个主 goroutine 和广播（broadcaster）goroutine，每一个连接里面又包含一个连接处理（handleConn）goroutine 和一个客户写入（clientwriter）goroutine
  广播器（broadcaster）是用于如何使用 select 的一个规范说明，因为它需要对三种不同的消息进行响应
  主 goroutine 的工作是监听端口，接受连接客户端的网络连接，对每一个连接，它将创建一个新的 handleConn goroutine
 */
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string  // 对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)  // 所有连接的客户端
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)  // 对外发送客户消息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "欢迎 " + who
	messages <- who + " 上线"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略 input.Err() 中可能的错误

	leaving <- ch
	messages <- who + " 下线"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)  // 注意：忽略网络层面的错误
	}
}