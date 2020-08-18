package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Go 语言中 sync 包里提供了互斥锁 Mutex 和读写锁 RWMutex 用于处理并发过程中可能出现同时两个或多个协程（或线程）读或写同一个变量的情况
func main() {
	mutex()
	rwMutex()
}

/*
  需要注意的是一个互斥锁 Mutex 只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
  互斥锁中其有两个方法可以调用：
    func (m *Mutex) Lock()
    func (m *Mutex) Unlock()
 */
func mutex() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second)
}

/*
  读写锁有如下四个方法：
    写操作的锁定和解锁分别是 func (*RWMutex) Lock 和 func (*RWMutex) Unlock
    读操作的锁定和解锁分别是 func (*RWMutex) Rlock 和 func (*RWMutex) RUnlock

  读写锁的区别在于：
    当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁
    当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续
    当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定

  我们可以将其总结为如下三条：
    同时只能有一个 goroutine 能够获得写锁定
    同时可以有任意多个 gorouinte 获得读锁定
    同时只能存在写锁定或读锁定（读和写互斥）
 */
func rwMutex() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

var count int
var rw sync.RWMutex

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}