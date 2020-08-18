package main

import (
	"context"
	"fmt"
	"time"
)

/*
  Context 在 Go 1.7 之后就加入到了 Go 语言标准库中，准确说它是 Goroutine 的上下文，包含 Goroutine 的运行状态、环境、现场等信息

  Context 包的核心就是 Context 接口，其定义如下：
	type Context interface {
      Deadline() (deadline time.Time, ok bool)
      Done() <-chan struct{}
      Err() error
      Value(key interface{}) interface{}
	}

  其中：
	Deadline 方法需要返回当前 Context 被取消的时间，也就是完成工作的截止时间（deadline）
	Done 方法需要返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个Channel
	Err 方法会返回当前 Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值：
		如果当前 Context 被取消就会返回 Canceled 错误
		如果当前 Context 超时就会返回 DeadlineExceeded 错误
	Value 方法会从 Context 中返回键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨 API 和进程间跟请求域的数据

  使用 Context 的注意事项：
    不要把 Context 放在结构体中，要以参数的方式显示传递
    以 Context 作为参数的函数方法，应该把 Context 作为第一个参数
    给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO
    Context 的 Value 相关方法应该传递请求域的必要数据，不应该用于传递可选参数
    Context 是线程安全的，可以放心的在多个 Goroutine 中传递

  Go 语言中的 Context 的主要作用还是在多个 Goroutine 或者模块之间同步取消信号或者截止日期，用于减少对资源的消耗和长时间占用，避免资源浪费，虽然传值也是它的功能之一，但是这个功能我们还是很少用到
 */
func main() {
	withCancel()
	withDeadline()
	withTimeout()
	withValue()
}

// WithCancel 返回带有新 Done 通道的父节点的副本，当调用返回的 cancel 函数或当关闭父上下文的 Done 通道时，将关闭返回上下文的 Done 通道
func withCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // return 结束该 goroutine，防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用 cancel，之后在上面 goroutine 的 select 中也会 return

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// WithDeadline 函数会返回父上下文的副本，并将 deadline 调整为不迟于 d。如果父上下文的 deadline 已经早于 d，则 WithDeadline(parent, d) 在语义上等同于父上下文。当截止日过期时，当调用返回的 cancel 函数时，或者当父上下文的 Done 通道关闭时，返回上下文的 Done 通道将被关闭，以最先发生的情况为准
func withDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管 ctx 会过期，但在任何情况下调用它的 cancel 函数都是很好的实践
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func withTimeout() {
	// 传递带有超时的上下文
	// 告诉阻塞函数在超时结束后应该放弃其工作。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 终端输出"context deadline exceeded"
	}
}

// WithValue 函数接收 context 并返回派生的 context，其中值 val 与 key 关联，并通过 context 树与 context 一起传递。这意味着一旦获得带有值的 context，从中派生的任何 context 都会获得此值。不建议使用 context 值传递关键参数，函数应接收签名中的那些值，使其显式化
func withValue() {
	type favContextKey string // 定义一个 key 类型

	// f: 一个从上下文中根据 key 取 value 的函数
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	// 创建一个携带 key 为 k，value 为 "Go" 的上下文
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favContextKey("color"))
}