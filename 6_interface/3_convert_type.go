package main

import "fmt"

/*
  接口和其他类型的转换可以在 Go 语言中自由进行，前提是已经完全实现
  接口断言类似于流程控制中的 if。但大量类型断言出现时，应使用更为高效的类型分支 switch 特性
 */
func main() {
	toInterface()
	toType()
}

func toInterface() {
	// 创建动物的名字到实例的映射，map[string]interface{} 表示 value 可以为任何类型
	animals := map[string]interface{} {
		"bird": new(bird),
		"pig":  new(pig),
	}

	// 遍历映射
	for name, obj := range animals {
		/*
		  类型断言：如果断言对象是断言指定的类型，则返回转换为断言对象类型的接口；如果不是指定的断言类型时，断言的第二个参数将返回 false
		*/
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}

		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}
}

// 定义飞行动物接口
type Flyer interface {
	Fly()
}

// 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 为鸟添加 Walk() 方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// 为猪添加 Walk() 方法, 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func toType() {
	p1 := new(pig)

	// 由于 pig 实现了 Walker 接口，因此可以被隐式转换为 Walker 接口类型保存于 a 中
	var a Walker = p1

	// 由于 a 中保存的本来就是 *pig 本体，因此可以转换为 *pig 类型
	p2 := a.(*pig)

	fmt.Printf("p1=%p p2=%p", p1, p2)
}