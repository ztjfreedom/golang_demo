package main

import "fmt"

/*
  Go 语言通过用自定义的方式形成新的类型，结构体是类型中带有成员的复合类型。Go 语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性
  Go 语言中的类型可以被实例化，使用 new 或 & 构造的类型实例的类型是类型的指针

  关于 Go 语言的类（class）
    Go 语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念
    Go 语言的结构体与“类”都是复合结构体，但 Go 语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性
    Go 语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法
 */
func main() {
	createInstance()
	initStruct()
	simulateConstructor()
}

func createInstance() {
	/*
	  方法 1：
	  结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化
	 */
	var p Point
	p.X = 10
	p.Y = 20
	fmt.Println(p)

	/*
	  方法 2：
	  Go 语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成"指针"类型的结构体
	  在 C/C++ 语言中，使用 new 实例化类型后，访问其成员变量时必须使用 -> 操作符
	  在 Go 语言中，访问结构体指针的成员变量时可以继续使用 . ，这是因为 Go 语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name
	 */
	tank := new(Tank)
	tank.Name = "Canon"
	tank.HealthPoint = 300
	fmt.Println(tank)

	/*
	  方法3：（使用最广泛）
	  在Go语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
	 */
	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	fmt.Println(cmd)

	// 函数封装
	cmd = newCommand("version", &version, "show version")
	fmt.Println(cmd)
}

// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装初始化过程
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

type Point struct {
	X int
	Y int
}

type Tank struct{
	Name string
	HealthPoint int
	MagicPoint int
}

type Command struct {
	Name    string    // 指令名称
	Var     *int      // 指令绑定的变量
	Comment string    // 指令的注释
}

func initStruct() {
	// 结构体可以使用“键值对”（Key value pair）初始化字段，每个“键”（Key）对应结构体中的一个字段，键的“值”（Value）对应字段需要初始化的值
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation)

	// Go 语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)

	// 初始化匿名结构体：匿名结构体没有类型名称，无须通过 type 关键字定义就可以直接使用
	msg := &struct {  // 定义部分
		id   int
		data string
	}{  // 值初始化部分
		1024,
		"hello",
	}
	fmt.Printf("%T\n", msg)
	fmt.Println(msg)
}

type People struct {
	name  string
	child *People  // 如果结构体成员中有同一类型的结构体的话，只能包含结构体的指针类型，包含非指针类型会引起编译错误
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func simulateConstructor() {
  // Go 语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数
  c1 := NewCatByName("Tom")
  c2 := NewCatByColor("Red")
  c3 := NewCat("Jim")
  c4 := NewBlackCat("Black")
  fmt.Println(c1, c2, c3, c4)
}

type Cat struct {
	Color string
	Name  string
}

// 模拟构造函数重载
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 模拟构造函数重载
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 模拟父级构造调用
type BlackCat struct {
	Cat  // 嵌入Cat, 类似于派生，使 BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员
}

// 模拟父级构造调用：构造基类
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 模拟父级构造调用：构造子类
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color  // BlackCat 中嵌入了 Cat，所以有 Color 成员
	return cat
}