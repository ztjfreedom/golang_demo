package main

import (
	"errors"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(div(1, 0))
	parse()

	/*
	  panic：有些错误只能在运行时检查，如数组访问越界、空指针引用等，这些运行时错误会引起宕机
	  一般而言，当宕机发生时，程序会中断运行，并立即执行在该 goroutine 中被延迟的函数（defer 机制）
	  由于 panic 会引起程序的崩溃，因此 panic 一般用于严重错误，如程序内部的逻辑不一致。任何崩溃都表明了我们的代码中可能存在漏洞，所以对于大部分漏洞，我们应该使用Go语言提供的错误机制，而不是 panic

	  Recover 是一个 G o语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行
	  在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过 try/catch 机制捕获异常，没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行
	  Go 语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常，recover 的宕机恢复机制就对应其他语言中的 try/catch 机制

	  panic 和 recover 的组合有如下特性：
	    有 panic 没 recover，程序宕机
	    有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行

	  虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性
	  在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃
	  如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置
	*/

	// 手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用 panic 传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		// 执行完对应的 defer 后，从宕机点退出当前函数后继续执行，所以这一行代码不会被打印
		fmt.Println("手动宕机后")
	})

	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})

	// 手动 panic
	defer fmt.Println("宕机后要做的事情 1")
	defer fmt.Println("宕机后要做的事情 2")
	panic("宕机")
}

/*
  Go 语言的错误处理思想及设计包含以下特征：
    一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误
    在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理

  Go 语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做，Go语言的设计者认为其他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源，同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程序将在不可预知的时刻崩溃
*/
// 自定义除数为 0 的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	// 判断除数为 0 的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

// 使用 errors.New 定义的错误字符串的错误类型是无法提供丰富的错误信息的，那么，如果需要携带错误信息返回，就需要借助自定义结构体实现错误接口
// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现 error 接口，返回错误描述
func (e *ParseError) Error() string {
	// Sprintf 格式化字符串
	return fmt.Sprintf("%s: %d", e.Filename, e.Line)
}

// 创建错误实例的函数
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func parse() {
	var e error

	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)

	// 通过error接口查看错误描述
	fmt.Println(e.Error())

	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError:
		// 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default:
		// 其他类型的错误
		fmt.Println("other error")
	}
}

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取 panic 传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error:
			// 运行时错误
			fmt.Println("runtime error:", err)
		default:
			// 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	entry()
}