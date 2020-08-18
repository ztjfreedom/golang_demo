package main

import "fmt"

func main() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%-3d", x, y, x*y)  // -3 表示 3 位左对齐
		}
		fmt.Println()
	}
}