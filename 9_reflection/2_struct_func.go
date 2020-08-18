package main

import (
	"fmt"
	"reflect"
)

func main() {
	reflectStructField()
	nilValid()
	canAddr()
	setVal()
	setStructVal()
	newInstance()
	callFunc()
}

// 定义结构体
type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

func reflectStructField() {
	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	// 获取字段数量
	fmt.Println("NumField", d.NumField())

	// 获取索引为2的字段 (float32 字段)
	floatField := d.Field(2)

	// 输出字段类型
	fmt.Println("Field", floatField.Type())

	// 根据名字查找字段
	fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())

	// 根据索引查找值中, next 字段的 int 字段的值
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
}

/*
  IsNil() bool:	返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的 v == nil 操作
  IsValid() bool: 判断值是否有效。当值本身非法时，返回 false，例如 reflect Value 不包含任何值，值为 nil 等
 */
func nilValid() {
	// *int 的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

	// nil 值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	// *int 类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	// 实例化一个结构体
	s := struct{}{}

	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())

	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())

	// 实例化一个 map
	m := map[int]int{}

	// 尝试从 map 中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

// 其中 a 对应的变量则不可取地址。因为 a 中的值仅仅是整数 2 的拷贝副本。b 中的值也同样不可取地址。c 中的值还是不可取地址，它只是一个指针 &x 的拷贝。实际上，所有通过 reflect.ValueOf(x) 返回的 reflect.Value 都是不可取地址的
// 但是对于 d，它是 c 的解引用方式生成的，指向另一个变量，因此是可取地址的。我们可以通过调用 reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的 Value
func canAddr() {
	x := 2  // value type variable
	a := reflect.ValueOf(2)  // 2 int no
	b := reflect.ValueOf(x)  // 2 int no
	c := reflect.ValueOf(&x)  // &x *int no
	d := c.Elem()  // 2 int yes (x)
	fmt.Println(a.CanAddr(), b.CanAddr(), c.CanAddr(), d.CanAddr())
}

// 值可修改条件之一：可被寻址
func setVal() {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量 a 的反射值对象 (a 的地址)
	valueOfA := reflect.ValueOf(&a)
	// 取出 a 地址的元素 (a 的值)
	valueOfA = valueOfA.Elem()
	// 修改 a 的值为 1
	valueOfA.SetInt(1)
	fmt.Println(valueOfA.Int())
}

// 值可修改条件之一：被导出
func setStructVal() {
	type dog struct {
		LegCount int  // LegCount 必须首字母大写，是被导出的字段
	}
	// 获取 dog 实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	// 取出 dog 实例地址的元素
	valueOfDog = valueOfDog.Elem()
	// 获取 legCount 字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")
	// 尝试设置 legCount 的值
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int())
}

// 通过类型信息创建实例
func newInstance() {
	var a int
	// 取变量 a 的反射类型对象
	typeOfA := reflect.TypeOf(a)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)
	// 输出 Value 的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())
}

func callFunc() {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
}

// 普通函数
func add(a, b int) int {
	return a + b
}