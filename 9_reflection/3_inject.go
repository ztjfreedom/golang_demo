package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
)

// inject 是依赖注入的Go语言实现，它能在运行时注入参数，调用方法，是 Martini 框架（Go语言中著名的 Web 框架）的基础核心
func main() {
	injFunc()
	injStruct()
	interfaceOf()
	parentTypeValid()
	parentInvoke()
	parentApply()
}

type S1 interface{}
type S2 interface{}
type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name=%s, company=%s, level=%s, age=%d!\n", name, company, level, age)
}

func injFunc() {
	// 控制实例的创建
	inj := inject.New()
	// 实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))  // S1, S2 是自定义 type
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	// 函数反转调用
	inj.Invoke(Format)
}

func injStruct() {
	// 创建被注入实例
	s := Staff{}
	// 控制实例的创建
	inj := inject.New()
	// 初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	// 实现对 struct 注入
	inj.Apply(&s)
	fmt.Printf("s=%v\n", s)
}

type SpecialString interface{}

// InterfaceOf 方法的参数必须是一个接口类型的指针，如果不是则引发 panic。InterfaceOf 方法的返回类型是 reflect.Type
func interfaceOf() {
	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
}

// SetParent 方法用于给某个 Injector 指定父 Injector。Get 方法通过 reflect.Type 从 injector 的 values 成员中取出对应的值，它可能会检查是否设置了 parent，直到找到或返回无效的值
func parentTypeValid() {
	inj := inject.New()
	inj.Map("C语言中文网")  // string
	inj.MapTo("Golang", (*SpecialString)(nil))  // SpecialString
	inj.Map(20)  // int
	fmt.Println("字符串是否有效？", inj.Get(reflect.TypeOf("Go语言入门教程")).IsValid())
	fmt.Println("特殊字符串是否有效？", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
	fmt.Println("int 是否有效？", inj.Get(reflect.TypeOf(18)).IsValid())
	fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())

	inj2 := inject.New()
	inj2.Map([]byte("test"))  // []byte
	inj.SetParent(inj2)
	fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
}


func Say(name string, gender SpecialString, age int) {
	fmt.Printf("My name is %s, gender is %s, age is %d!\n", name, gender, age)
}

func parentInvoke() {
	inj := inject.New()
	inj.Map("张三")
	/*
	  MapTo 方法有一个额外的参数可以指定特定的类型当键，第二个参数 ifacePtr 必须是接口指针类型
	  为什么需要有 MapTo 方法？因为注入的参数是存储在一个以类型为键的 map 中，可想而知，当一个函数中有一个以上的参数的类型是一样时，后执行 Map 进行注入的参数将会覆盖前一个通过 Map 注入的参数
	 */
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(25)
	inj.SetParent(inj2)
	inj.Invoke(Say)
}

type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString `inject`
	uid    int           `inject`
	Age    int           `inject`
}

// Apply 方法是用于对 struct 的字段进行注入
func parentApply() {
	s := TestStruct{}
	inj := inject.New()
	inj.Map("张三")
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(26)
	inj.SetParent(inj2)
	inj.Apply(&s)
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
}