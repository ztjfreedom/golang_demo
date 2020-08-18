package main

import (
	"fmt"
	"math/big"
	"time"
)

/*
  Go 语言中 math/big 包实现了大数字的多精度计算，支持 Int（有符号整数）、Rat（有理数）和 Float（浮点数）等数字类型
 */
func main() {
	// SetUint64
	big1 := new(big.Int).SetUint64(uint64(1000))
	fmt.Println("big1 is: ", big1)
	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)

	// SetString
	big1, _ = new(big.Int).SetString("1000", 10)
	fmt.Println("big1 is: ", big1)
	big2 = big1.Uint64()
	fmt.Println("big2 is: ", big2)

	// fibonacci
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		if i == LIM - 1 {
			fmt.Printf("数列第 %d 位: %d\n", i+1, result)
		}
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为: %s\n", delta)
}

const LIM = 1000 // 求第 1000 位的斐波那契数列
var fibs [LIM]*big.Int // 使用数组保存计算出来的数列的指针，for dp

func fibonacci(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}