package main

import "fmt"

func main() {
	embed()
	namedEmbed()
	simulateInherit()
	initEmbed()
}

/*
  结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体
  可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在Go语言中，相比较于继承，组合更受青睐
 */
func embed() {
	// 使用结构体指针
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5  // 匿名内嵌的结构体可以直接访问其成员变量
	outer.in2 = 10
	fmt.Println("outer is:", outer)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int // anonymous field，在一个结构体中对于每一种数据类型只能有一个匿名字段
	innerS //anonymous field，内嵌结构体，也可以来自其他包
}

func namedEmbed() {
	outer := &outerStruct{}
	outer.b = 1
	outer.c = 2.5
	outer.in.in1 = 10  // 这里不是匿名的，所以需要 .in
	outer.in.in2 = 20
	fmt.Println("outer is:", *outer)
}

type outerStruct struct {
	b int
	c float32
	in innerS
}

/*
  模拟继承：
  Go 语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性
 */
func simulateInherit() {
	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	// 实例化人类
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}

// 可飞行的
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("can walk")
}

// 人类
type Human struct {
	Walkable // 人类能行走
}

// 鸟类
type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类能飞行
}

func initEmbed() {
	normal := NormalCar {
		Wheel: Wheel {
			Size: 18,
		},
		Engine: Engine {
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", normal)

	sport := &SportCar {
		Wheel: Wheel {
			Size: 20,
		},
		Engine: struct {  // 初始化被嵌入的结构体时，就需要再次声明结构才能赋予数据
			Power int
			Type  string
		}{
			Type:  "3.0T",
			Power: 400,
		},
	}
	fmt.Printf("%+v\n", sport)
}

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int     // 功率
	Type  string  // 类型
}

// 车
type NormalCar struct {
	Wheel  // 匿名
	Engine
}

type SportCar struct {
	Wheel
	// 有时考虑编写代码的便利性，会将结构体直接定义在嵌入的结构体中，初始化被嵌入的结构体时，就需要再次声明结构才能赋予数据
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}