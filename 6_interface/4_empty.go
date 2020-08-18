package main

import "fmt"

/*
  空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值
 */
func main() {
	setValue()
	getValue()
	compare()
	dict()
}

func setValue() {
	var any interface{}

	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)
}

func getValue() {
	var a int = 100
	var i interface{} = a

	// i 在赋值完成后的内部值为 int，但 i 还是一个 interface{} 类型的变量，需要使用类型断言转换
	var b int = i.(int)
	fmt.Println(b)
}

func compare() {
	var a interface{} = 100
	var b interface{} = "hi"

	// 当接口中保存有动态类型的值时（比如切片），不能进行比较，否则会触发 panic
	fmt.Println(a == b)
}

func dict() {
	// 创建字典实例
	dict := NewDictionary()

	// 添加游戏数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	// 遍历所有的字典元素
	dict.Visit(func(key, value interface{}) bool {
		// 将值转为 int 类型，并判断是否大于 40
		if value.(int) > 40 {
			// 输出很贵
			fmt.Println(key, "is expensive")
			return true
		}
		// 默认都是输出很便宜
		fmt.Println(key, "is cheap")
		return true
	})
}

// 字典结构
type Dictionary struct {
	data map[interface{}]interface{}  // 键值都为 interface{} 类型，可以将任意类型的值做成键值对保存
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// 设置键值
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值，如果回调返回值为 false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 清空所有的数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}