package main

import (
	"fmt"
	"sync"
)

func main() {
	waitChan()
	waitGroup()
}

func waitChan() {
	done := make(chan int, 5) // 带 5 个缓存

	// 开 N 个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("C语言中文网")
			done <- 1
		}()
	}

	// 等待 N 个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func waitGroup() {
	var wg sync.WaitGroup

	// 开 5 个后台打印线程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Golang")
			wg.Done()
		}()
	}

	// 等待 N 个后台线程完成
	wg.Wait()
}