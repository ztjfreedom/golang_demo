package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU() // 获得当前设备的 cpu 核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) // 设置需要用到的 cpu 数量

	task(cpuNum)
}

func task(cpuNum int) {
	c := make(chan int, cpuNum)
	result := make([]int, cpuNum)
	for i := 0; i < cpuNum; i++ {
		go subTask(i, result, c)
	}

	for i := 0; i < cpuNum; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	fmt.Println(result)
}

func subTask(i int, result []int, c chan int) {
	for j := 0; j < 100000; j ++ {
		result[i] ++
	}
	c <- 1  // 发信号告诉任务管理者我已经计算完成了
}