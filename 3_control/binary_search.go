package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	binarySearch(&arr, 0, len(arr) - 1, 30)
	fmt.Println("main arr =",arr)
}


func binarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int)  {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := leftIndex + (rightIndex - leftIndex) / 2

	if (*arr)[middle] > findVal {
		binarySearch(arr, leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal {
		binarySearch(arr, middle + 1, rightIndex, findVal)
	} else {
		fmt.Println("找到了，下标为：", middle)
	}
}