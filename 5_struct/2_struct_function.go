package main

import (
	"fmt"
)

func main() {
	b := new(Bag)
	b.Insert(1)
	fmt.Println(b)

	p := &Property{}
	p.SetValue(100)
	fmt.Println(p.Value())

	p1 := Position{1, 1}
	p2 := Position{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)

	p1.setX(100)
	fmt.Println(p1.X)
}

type Bag struct {
	items []int
}

/*
  结构体的方法：
  (b *Bag) 是接收器
 */
func (b *Bag) Insert(itemId int) {
	b.items = append(b.items, itemId)
}

type Property struct {
	value int
}

/*
  指针类型的接收器：
  指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this 或者 self
  由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
 */
func (p *Property) SetValue(v int) {
	// 修改p的成员变量
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

type Position struct {
	X int
	Y int
}
/*
  非指针类型的接收器：
  当方法作用于非指针接收器时，Go 语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效
 */
func (p Position) Add(other Position) Position {
	// 成员值与参数相加后返回新的结构
	return Position{p.X + other.X, p.Y + other.Y}
}

func (p Position) setX(x int) {
	// 非指针接收器的方法中可以获取接收器的成员值，但修改后无效
	p.X = x
}