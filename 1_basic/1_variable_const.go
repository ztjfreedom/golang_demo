package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unicode"
)

/*
  全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，不包含这个全局变量的源文件需要使用 import 关键字引入全局变量所在的源文件之后才能使用
  全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写
*/
const PI float32 = 3.14

const (
	NAME = "brebro"
	AGE = 24
)

func main() {
	declare()
	initial()
	assign()
	floatType()
	complexType()
	strType()
	byteRuneType()
	changeType()
	constType()
	enum()
	alias()
	convertNumString()
}

func declare() {
	/*
	标准格式：
	var 变量名 变量类型
	var a, b *int

	批量格式：
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	*/

	// 简短格式：
	x := 100
	i, j := 0, 1
	a, s := 1, "abc"
	fmt.Println(x, i, j, a, s);
}

func initial() {
	/*
	标准格式
	var 变量名 类型 = 表达式
	var hp int = 100

	编译器推导类型格式
	var hp = 100

	简短格式
	hp := 100
	conn, err := net.Dial("tcp","127.0.0.1:8080")
	*/

	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack - defence) * damageRate
	fmt.Println(damage)
}

func assign() {
	// 多变量同时赋值
	var a int = 100
	var b int = 200
	b, a = a, b

	// 匿名变量
	c, _ := getData()
	_, d := getData()

	fmt.Println(a, b, c, d)

	fmt.Println(sum(a, b))
}

func getData() (int, int) {
	return 1000, 2000
}

func sum(a, b int) int {
	// a, b 是形参
	num := a + b
	return num
}

func floatType() {
	a := .7182
	b := 1.
	c := 6.02e5
	d := 6.62e-5
	fmt.Printf("%f %.3f %.f %.6f\n", a, b, c, d)
}

func complexType() {
	var x complex128 = complex(1, 2)  // 1 + 2i
	var y complex128 = complex(3, 4)  // 3 + 4i
	var z complex128 = complex(3, 4)  // 3 + 4i
	fmt.Println(x*y, real(x*y), imag(x*y), y == z)  // "(-5 + 10i), -5, 10, true"
}

func strType() {
	s := "Hel" + "lo,\t"
	s += "World!"
	const str = `第一行
第二行
\r\n`
	fmt.Println(s, string(s[0]))
	fmt.Println(str)
}

func byteRuneType() {
	var ch int = '\u0041'
	fmt.Printf("%d - %c - %X - %U\n", ch, ch, ch, ch)  // integer, character, UTF-8 bytes, UTF-8 code point
	s := "A 0"
	fmt.Println(unicode.IsDigit(rune(s[2])), unicode.IsLetter(rune(s[0])), unicode.IsSpace(rune(s[1])), unicode.IsDigit(rune(s[0])))
}

func changeType() {
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16, "int32 range:", math.MinInt32, math.MaxInt32)
	var a int32 = 1047483647
	b := int16(a)
	var c float32 = math.Pi
	fmt.Println(a, b, int(c))
}

func constType() {
	/*
	  常量是在编译时被创建的，即使定义在函数内部也是如此，并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型
	  因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度
	 */
	const LEN = 2
	var arr [LEN]int
	arr[0] = 1
	arr[1] = 2

	// 定义常量的表达式必须为能被编译器求值的常量表达式
	const NUM = 2/3

	// 批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(arr, NUM, a, b, c, d)

	// iota 常量生成器
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday, Friday)

	/*
	  math.Pi 是无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方
 	  有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串
	*/
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}

func enum() {
	// Go语言现阶段没有枚举类型，但是可以使用 const 常量配合 iota 来模拟枚举类型
	type Weapon int
	const (
		Arrow Weapon = iota    // 开始生成枚举值, 默认为 0
		Rifle  // 默认类型和上面一样，就是 Weapon
		Blower
	)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(Arrow, Rifle, Blower, weapon)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)

	// 将枚举值转换为字符串
	fmt.Printf("%s %d\n", CPU, CPU)
}

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {  // 定义 ChipType 类型的方法 String()，返回值为字符串类型，当这个类型需要显示为字符串时，Go 语言会自动寻找 String() 方法并进行调用
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func alias() {
	// 将 NewInt 定义为 int 类型
	type NewInt int
	// 将 int 取一个别名叫 IntAlias
	type IntAlias = int

	var a NewInt
	var b IntAlias
	fmt.Printf("%T %T\n", a, b)

	// 声明变量 v 为车辆类型
	var v Vehicle
	// 指定调用 FakeBrand 的 Show
	v.FakeBrand.Show()
	// 取 v 的类型反射对象
	tv := reflect.TypeOf(v)
	// 遍历 v 的所有成员
	for i := 0; i < tv.NumField(); i++ {
		// v 的成员信息
		f := tv.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加 Show() 方法
func (t Brand) Show() {
}

// 为 Brand 定义一个别名 FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构，其中一个是别名
	FakeBrand
	Brand
}

func convertNumString() {
	// Itoa()：整型转字符串
	fmt.Printf("type:%T value:%#v\n", strconv.Itoa(100), strconv.Itoa(100))

	// Atoi()：字符串转整型
	num, err := strconv.Atoi("100")
	if err != nil {
		fmt.Printf("转换失败！")
	} else {
		fmt.Printf("type:%T value:%#v\n", num, num)
	}

	// Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()
	// ParseBool() 函数用于将字符串转换为 bool 类型的值，它只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误
	boo, err := strconv.ParseBool("t")
	if err != nil {
		fmt.Printf("转换失败！")
	} else {
		fmt.Println(boo)
	}

	// ParseInt()
	// base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
	// bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64
	// ParseUint() 函数的功能类似于 ParseInt() 函数，但 ParseUint() 函数不接受正负号，用于无符号整型
	num64, err := strconv.ParseInt("-11", 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num64)
	}

	// ParseFloat()：如果 s 合乎语法规则，函数会返回最为接近 s 表示值的一个浮点数
	// bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64
	flo, err := strconv.ParseFloat("3.1415926", 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(flo)
	}

	// Format 系列函数实现了将给定类型数据格式化为字符串类型的功能，其中包括 FormatBool()、FormatInt()、FormatUint()、FormatFloat()
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(100, 10))  // base 代表几进制，FormatUint() 函数与 FormatInt() 函数的功能类似，但是参数 i 必须是无符号的 uint64 类型

	/*
	  bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入
	  fmt 表示格式，可以设置为 “f” 表示 -ddd.dddd、“b” 表示 -ddddp±ddd，指数为二进制、“e” 表示 -d.dddde±dd 十进制指数、“E” 表示 -d.ddddE±dd 十进制指数、“g” 表示指数很大时用 “e” 格式，否则 “f” 格式、“G” 表示指数很大时用 “E” 格式，否则 “f” 格式
	  prec 控制精度（排除指数部分）：当参数 fmt 为 “f”、“e”、“E” 时，它表示小数点后的数字个数；当参数 fmt 为 “g”、“G” 时，它控制总的数字个数。如果 prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f
	 */
	fmt.Println(strconv.FormatFloat(3.1415926, 'E', -1, 64))

	// Append 系列函数用于将指定类型转换成字符串后追加到一个切片中，其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint()
	// 声明一个 slice
	b10 := []byte("int(base 10): ")
	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
}