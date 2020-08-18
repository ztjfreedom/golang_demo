package main

import (
	"fmt"
)

func main() {
	Arr := []int{23, 65, 13, 27, 42, 15, 38, 21, 4, 10}
	qSort(Arr, 0, len(Arr)-1)
	fmt.Println(Arr)
}

/*
  快速排序：分治法+递归实现
  随意取一个值A，将比A大的放在A的右边，比A小的放在A的左边；然后在左边的值AA中再取一个值B，将AA中比B小的值放在B的左边，将比B大的值放在B的右边。以此类推
*/
func qSort(arr []int, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}

	// 将大于 arr[flag] 的都放在右边，小于的，都放在左边
	for first < last {
		// 如果 flag 从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比 flag 小的
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}

	qSort(arr, left, flag - 1)
	qSort(arr, flag + 1, right)
}