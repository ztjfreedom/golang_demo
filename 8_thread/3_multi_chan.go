package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 5; i ++ {
			ch1 <- 1
			time.Sleep(time.Second)
		}
		// close(ch1)
	}()
	go func() {
		for i := 0; i < 5; i ++ {
			ch2 <- 2
			time.Sleep(time.Second)
		}
		// close(ch2)
	}()

	for {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println(data)
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println(data)
			}
		case <-time.After(time.Second * 3):
			fmt.Println("没有新的数据，退出程序")
			return
		}
	}
}