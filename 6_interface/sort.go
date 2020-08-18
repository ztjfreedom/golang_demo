package main

import (
	"fmt"
	"sort"
)

func main() {
	stringSort()
	intSort()
	oldCustomTypeSort()
	newCustomTypeSort()
}

func stringSort() {
	names := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Strings(names)
	fmt.Println(names)
}

func intSort() {
	ages := []int {5, 4, 3, 2, 1}
	sort.Ints(ages)
	fmt.Println(ages)
}

/*
  Go 语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片
  一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法
  type Interface interface {
      Len() int            // 获取元素数量
      Less(i, j int) bool // i，j是序列元素的指数。
      Swap(i, j int)        // 交换元素
  }
*/
func oldCustomTypeSort() {
	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	// 使用sort包进行排序
	sort.Sort(names)

	fmt.Println(names)
}

// 将 []string 定义为 MyStringList 类型
type MyStringList []string

// 实现 sort.Interface 接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现 sort.Interface 接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现 sort.Interface 接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

/*
  从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法
 */
func newCustomTypeSort() {
	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v ", v)
	}
	fmt.Println()
}

type HeroKind int

const (
	None = iota
	Tank
	Assassin
	Mage
)
type Hero struct {
	Name string
	Kind HeroKind
}