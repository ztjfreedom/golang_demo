package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeOf()
	kind()
	kindPtr()
	reflectStruct()
	reflectRule()
	setValue()
	setStruct()
	getValue()
}

// 在Go语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
func typeOf() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

/*
  在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便
  Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr
 */
func kind() {
	// 声明一个空结构体
	type cat struct {}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 获取 Zero 常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

// Go 语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
func kindPtr() {
	// 声明一个空结构体
	type cat struct {}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())

	// 取类型的元素
	// Go 语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个 * 操作
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

/*
  任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息
  FieldByName, FieldByIndex, FieldByNameFunc
 */
func reflectStruct() {
	// 声明一个空结构体
	type cat struct {
		Name string
		// 带有结构体 tag 的字段，JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从 tag 中取出需要的 tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

func reflectRule() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var y uint8 = 'x'
	vy := reflect.ValueOf(y)
	fmt.Println("type:", vy.Type())                            // uint8.
	fmt.Println("kind is uint8: ", vy.Kind() == reflect.Uint8) // true.
	y = uint8(vy.Uint())                                       // v.Uint returns a uint64.
	fmt.Println(y)
}

/*
  如果要修改“反射类型对象”其值必须是“可写的”
  “可写性”有些类似于寻址能力，但是更严格，它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。“可写性”最终是由一个反射对象是否存储了原始值而决定的
 */
func setValue() {
	var x float64 = 3.4
	q := reflect.ValueOf(x)
	fmt.Println("settability of q:", q.CanSet())

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(x, v.Interface())  // Interface returns v's current value as an interface{}
}

// 我们一般使用反射修改结构体的字段，只要有结构体的指针，我们就可以修改它的字段
func setStruct() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

/*
  可以通过下面几种方法从反射值对象 reflect.Value 中获取原值
  Interface(), Int(), Uint(), Float(), Bool(), Bytes(), String()
 */
func getValue() {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量 a 的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取 interface{} 类型的值, 通过类型断言转换
	var getA int = valueOfA.Interface().(int)
	// 获取 64 位的值, 强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, getA2)
}