package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
  在运行程序时，为运行参数加入-race参数，开启运行时（runtime）对竞态问题的分析，命令如下：
  go run -race race_detect.go
*/
func main() {
	// 生成 10 个并发序列号
	for i := 0; i < 10; i++ {
		go GenID()
	}

	time.Sleep(time.Second)
	fmt.Println(GenID())
}

var (
	// 序列号
	seq int64
)

// 序列号生成器
// func GenID() int64 {
//	 // 尝试原子的增加序列号
//	 atomic.AddInt64(&seq, 1)
//	 return seq
// }

func GenID() int64 {
	// 尝试原子的增加序列号
	return atomic.AddInt64(&seq, 1)
}