package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
  goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心
  goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，而且Go语言内部也实现了 goroutine 之间的内存共享
  使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，这样该函数执行时便会作为一个独立的并发线程，这种线程在 Go 语言中则被称为 goroutine

  所有 goroutine 在 main() 函数结束时会一同结束
  终止 goroutine 的最好方法就是自然返回 goroutine 对应的函数
 */

func main() {
	// Go 1.5 版本之前，默认使用的是单核心执行。从 Go 1.5 版本开始，默认执行上面语句以便让代码并发执行，最大效率地利用 CPU
	// Go 语言在 GOMAXPROCS 数量与任务数量相等时，可以做到并行执行，但一般情况下都是并发执行
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	// normalFuncGoRoutine()
	//anonymousFuncGoRoutine()
	//atomicAdd()
	//atomicLoadStore()
	mutexLock()
}

// 代码执行后，命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行
// 命名函数 goroutine
func normalFuncGoRoutine() {
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	// 构建一个无限循环
	for {
		times ++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

// 匿名函数 goroutine
func anonymousFuncGoRoutine() {
	go func() {
		var times int
		for {
			times ++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	var input string
	fmt.Scanln(&input)
}

// Go 语言可以使用 channel 处理数据贡献，但也提供了传统的方式
func atomicAdd() {
	wg.Add(2)
	go incCounter()
	go incCounter()
	wg.Wait() //等待 goroutine 结束
	fmt.Println(counter)
}


var (
	shutdown int64
	counter  int64
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func incCounter() {
	defer wg.Done()
	for count := 0; count < 2; count ++ {
		atomic.AddInt64(&counter, 1)  // 安全地对 counter 加 1
		runtime.Gosched()  // Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
	}
}

func atomicLoadStore() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

func mutexLock() {
	wg.Add(2)
	go incCounterLock(1)
	go incCounterLock(2)
	wg.Wait()
	fmt.Println(counter)
}

func incCounterLock(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个 goroutine 进入这个临界区
		mutex.Lock()

		value := counter
		runtime.Gosched()
		value ++
		counter = value

		mutex.Unlock() // 释放锁，允许其他正在等待的 goroutine 进入临界区
	}
}