package main

import "fmt"

/*
  接口本身是调用方和实现方均需要遵守的一种协议，大家按照统一的方法命名参数类型和数量来协调逻辑处理的过程
  Go 语言中使用组合实现对象特性的描述。对象的内部使用结构体内嵌组合对象应该具有的特性，对外通过接口暴露能使用的特性
  Go 语言的接口设计是非侵入式的，接口编写者无须知道接口被哪些类型实现。而接口实现者只需知道实现的是什么样子的接口，但无须指明实现哪一个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现
 */
func main() {
	implInterface()
	typeAssert()
	getType(10)
}

/*
  Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念
  但是 Go 语言里有非常灵活的接口概念，通过它可以实现很多面向对象的特性。很多面向对象的语言都有相似的接口概念，但Go语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了
  这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当我们使用的类型来自于不受我们控制的包时这种设计尤其有用
  type 接口类型名 interface {
      方法名1( 参数列表1 ) 返回值列表1
      方法名2( 参数列表2 ) 返回值列表2
      …
  }
 */
func implInterface() {
	// 实例化file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 将接口赋值 f，也就是 *file 类型
	writer = f

	// 使用DataWriter接口进行数据写入
	writer.WriteData("hello")
}

// 定义接口
type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

// 定义文件结构，用于实现 DataWriter
type file struct {
}

/*
  实现接口的条件：接口的方法与实现接口的类型方法格式一致
  在类型中添加与接口签名一致的方法就可以实现该方法。签名包括方法中的名称、参数列表、返回参数列表。也就是说，只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，那么接口的这个方法就不会被实现
  当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用

  Go 语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计
  对于 Go 语言来说，非侵入式设计让实现者的所有类型均是平行的、组合的。如何组合则留到使用者编译时再确认。因此，使用GO语言时，不需要同时也不可能有“类派生图”，开发者唯一需要关注的就是“我需要什么？”，以及“我能实现什么？”
 */
// 实现 DataWriter 接口的 WriteData 方法：给类型 file 添加与接口签名一致的方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func (d *file) CanWrite() bool {
	return true
}

/*
  一个类型可以实现多个接口
  多个类型可以实现同一个接口
 */
type Product interface {
	CanSell() bool
}

type Phone interface {
	CanCall() bool
}

type iphone struct {
}

func (i *iphone) CanSell() bool {
	return true
}

func (i *iphone) CanCall() bool {
	return true
}

type cloth struct {
}

func (c *cloth) CanSell() bool {
	return true
}

/*
  类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型
  在Go语言中类型断言的语法格式如下：
    value, ok := x.(T)
  其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）
 */
func typeAssert() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ok)

	x = "Hello"
	value, ok = x.(int)
	fmt.Println(value, ok)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type is int")
	case string:
		fmt.Println("the type is string")
	case float64:
		fmt.Println("the type is float")
	default:
		fmt.Println("unknown type")
	}
}