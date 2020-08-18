package main

import (
	"fmt"
	"sync"
)

/*
  Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex
  RWMutex 相对友好些，是经典的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine 可同时获取读锁（调用 RLock() 方法；而写锁（调用 Lock() 方法）会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占
 */
func main() {
	syncMutex()
	syncRWMutex()
}

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
	rwCountGuard sync.RWMutex
)

func syncMutex() {
	// 可以进行并发安全的设置
	SetCount(1)

	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}

func GetCount() int {
	// 锁定
	countGuard.Lock()

	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func syncRWMutex() {
	SetCountRW(1)
	fmt.Println(GetCountRW())
}

func GetCountRW() int {
	// 锁定
	rwCountGuard.RLock()

	// 在函数退出时解除锁定
	defer rwCountGuard.RUnlock()

	return count
}

func SetCountRW(c int) {
	rwCountGuard.Lock()
	count = c
	rwCountGuard.Unlock()
}