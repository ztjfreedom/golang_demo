package main

import "fmt"

func main() {
	// 结构的方法与普通函数
	delegate()

	// 实例化一个角色
	a := new(Actor)
	// 注册名为 OnSkill 的回调
	RegisterEvent("OnSkill", a.OnEvent)
	// 再次在 OnSkill 上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)
	// 调用事件，所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)
}

// Go 语言可以将类型的方法与普通函数视为一个概念，从而简化方法和函数混合作为回调类型时的复杂性。这个特性和 C# 中的代理（delegate）类似，调用者无须关心谁来支持调用，系统会自动处理是否调用普通函数或类型的方法
type class struct {
}

// 给结构体添加 Do 方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

// 普通函数的 Do
func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func delegate() {
	// 声明一个函数回调
	var delegate func(int)

	// 创建结构体实例
	c := new(class)

	// 将回调设为 c 的 Do 方法
	delegate = c.Do

	// 调用
	delegate(100)

	// 将回调设为普通函数
	delegate = funcDo

	// 调用
	delegate(100)
}

/*
  一个事件系统拥有如下特性：
    能够实现事件的一方，可以根据事件 ID 或名字注册对应的事件
    事件发起者，会根据注册信息通知这些注册者
    一个事件可以有多个实现方响应
 */
// 实例化一个通过字符串映射函数切片的 map
var eventByName = make(map[string][]func(interface{}))

// 注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]
	// 在列表切片中添加函数
	list = append(list, callback)
	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]
	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

// 声明角色的结构体
type Actor struct {
}

// 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}