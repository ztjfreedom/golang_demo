package main

import "fmt"

func main() {
	array()
	highDim()
	sli()
	appendSlice()
	copySlice()
	delFromSlice()
	rangeSlice()
	highDimSlice()
}

func array() {
	var a [3]int
	var b [3]string
	var c [3]int = [3]int{1, 2, 3}
	var d [3]int = [3]int{1, 2}  // 未指定的元素初始化为 0
	e := [...]int{1, 2, 3}  // 数组的长度是根据初始化值的个数来计算
	fmt.Println(a[2], b[2], c[2], d[2], e[2])

	// 可以直接通过较运算符 == 和 != 来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的
	fmt.Println(a == c, c == e)

	var team [3]string = [3]string {"a", "b", "c"}
	for k, v := range team {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()
}

func highDim() {
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var a [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	a = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中指定的元素
	a = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(a, a[1], a[1][0])

	var b [2]int = a[1]
	a[2] = b
	fmt.Println(a)
}

func sli() {
	// 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型），这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2], a[:], a[0:0])

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片（已经被分配了内存）
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果，切片是动态结构，只能与 nil 判定相等，不能互相判定相等。声明新的切片后，可以使用 append() 函数向切片中添加元素
	fmt.Println(strList == nil, numList == nil, numListEmpty == nil)

	strList = append(strList, "a")
	numList = append(numList, 1)
	numListEmpty = append(numListEmpty, 2)
	fmt.Println(strList, numList, numListEmpty)

	// 如果需要动态地创建一个切片，可以使用 make() 内建函数
	// 使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作
	b := make([]int, 2)
	c := make([]int, 2, 10)  // 其中 b 和 c 均是预分配 2 个元素的切片，只是 b 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素
	fmt.Println(b, c)
	fmt.Println(len(b), len(c))
}

func appendSlice() {
	var a []int
	a = append(a, 1) // 追加 1 个元素
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)

	var b = []int{1, 2, 3}
	b = append([]int{0}, b...) // 在开头添加 1 个元素
	b = append([]int{-3,-2,-1}, b...) // 在开头添加 1 个切片
	fmt.Println(b)

	// 因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来
	c := []int{100, 200, 300}
	c = append(c[:1], append([]int{150}, c[1:]...)...) // 在第 1 个位置插入 150
	c = append(c[:3], append([]int{220, 250}, c[3:]...)...) // 在第 3 个位置插入切片
	fmt.Println(c)
}

func copySlice() {
	// copy() 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据，切片不会因为等号操作进行元素的复制
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount - 1])
	// 复制原始数据从 4 到 6 (不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	fmt.Println()
}

func delFromSlice() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[1:], a[:len(a) - 1])  // 删除头尾

	b := append(a[:2], a[3:]...)  // 删除中间
	fmt.Println(b)
}

func rangeSlice() {
	/*
	  range 以配合关键字 for 来迭代切片里的每一个元素
	  range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	*/
	slice := []int{10, 20}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d ValueAddr: %X ElemAddr: %X\n", index, value, &value, &slice[index])
	}

	slice = []int{10, 20, 30, 40}
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}

func highDimSlice() {
	// slice := [][]int{{10}, {100, 200}}
	var slice [][]int
	slice = [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
	fmt.Println(slice)
}