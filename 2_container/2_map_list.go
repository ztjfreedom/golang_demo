package main

import (
	"container/list"
	"fmt"
	"sort"
	"sync"
)

func main() {
	createMap()
	traverseMap()
	delFromMap()
	multiKey()
	syncMap()
	createList()
}

func createMap() {
	// 未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}

	/*
	  mapCreated := make(map[string]float) 等价于 mapCreated := map[string]float{}
	  注意：可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	  make 和 new 关键字的主要区别:
	  make 关键字的主要作用是创建 slice、map 和 Channel 等内置的数据结构，而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针
	 */
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Println("mapLit values:", mapLit["one"], mapLit["two"])
	fmt.Println("mapAssigned values:", mapLit["one"], mapLit["two"])
	fmt.Println("mapCreated values:", mapLit["key1"], mapLit["key2"])

	// value 为 slice
	mp := make(map[int][]int)
	mp[1] = []int{100, 200}
	fmt.Println("mp values:", mp[1])
}

func traverseMap() {
	scene := map[string]int{}
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()

	// 如果需要特定顺序的遍历结果，正确的做法是先排序
	// 声明一个切片保存 map 数据
	var sceneList []string
	// 将 map 数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	for _, k := range sceneList {
		fmt.Printf("%s:%d ", k, scene[k])
	}
	fmt.Println()
}

func delFromMap() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()

	// Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多
	scene = make(map[string]int)
	fmt.Println(len(scene))
}

func multiKey() {
	/*
	  Go 语言的底层会为 map 的键自动构建哈希值。能够构建哈希值的类型必须是非动态类型、非指针、函数、闭包
	    非动态类型：可用数组，不能用切片
	    非指针：每个指针数值都不同，失去哈希意义
	    函数、闭包不能作为 map 的键
	 */
	list := []Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}

	buildIndex(list)
	queryData("张三", 30)
}

// 人员档案
type Profile struct {
	Name    string   // 名字
	Age     int      // 年龄
	Married bool     // 已婚
}

// 查询键
type queryKey struct {
	Name string
	Age  int
}

// 创建查询键到数据的映射
var mapper = make(map[queryKey]Profile)

// 构建查询索引
func buildIndex(list []Profile) {
	// 遍历所有数据
	for _, profile := range list {
		// 构建查询键
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		// 保存查询键
		mapper[key] = profile
	}
}

// 根据条件查询数据
func queryData(name string, age int) {
	// 根据查询条件构建查询键
	key := queryKey{name, age}
	// 根据键值查询数据
	result, ok := mapper[key]
	// 找到数据打印出来
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}

func syncMap() {
	/*
	  Go 语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的
	  Go 语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构
	  sync.Map 有以下特性：
	    无须初始化，直接声明即可
	    sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除
	    使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false
	 */

	// 不需要初始化
	var scene sync.Map

	// 将键值对保存到 sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从 sync.Map 中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有 sync.Map 中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Printf("%s:%d ", k, v)
		return true
	})
	fmt.Println()
}

func createList() {
	/*
	  列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等
	  在 Go 语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作
	  列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机
	 */
	l := list.New()
	// 尾部添加 canon
	l.PushBack("canon")
	// 头部添加 67 - canon
	l.PushFront(67)
	// 尾部添加后保存元素句柄 67 - canon - fist
	element := l.PushBack("fist")
	// 在 fist 之后添加 high 67 - canon - fist - high
	l.InsertAfter("high", element)
	// 在 fist 之前添加 noon 67 - canon - noon - fist - high
	l.InsertBefore("noon", element)
	// 删除 fist 67 - canon - noon - high
	l.Remove(element)

	// 遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ")
	}
	fmt.Println()
}