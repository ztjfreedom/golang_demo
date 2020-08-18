package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
  Go 语言中除了可以使用通道（channel）和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的同步，等待组可以保证在并发环境中完成指定数量的任务
  在 sync.WaitGroup（等待组）类型中，每个 sync.WaitGroup 值在内部维护着一个计数，此计数的初始默认值为零
 */
func main() {
	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {

			// 使用defer, 表示函数完成时将等待组值减 1，等价于 Add(-1)
			defer wg.Done()

			// 使用http访问提供的地址
			_, err := http.Get(url)

			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)

			// 通过参数传递url地址
		}(url)
	}

	// 等待所有的任务完成，当等待组计数器不等于 0 时阻塞直到变 0
	wg.Wait()

	fmt.Println("over")
}