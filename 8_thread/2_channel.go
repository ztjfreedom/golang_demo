package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
  channels 是一个通信机制，它可以让一个 goroutine 通过它给另一个 goroutine 发送值信息。每个 channel 都有一个特殊的类型
  Go 语言提倡使用通信的方法代替共享内存，当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了确保同步交换数据的机制。声明通道时，需要指定将要被共享的数据的类型。可以通过通道共享内置类型、命名类型、结构类型和引用类型的值或者指针

  Go 语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信
  通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序
 */
func main() {
	// sendReceive()
	// loopReceive()
	// unidirectionalChan()
	// unbufferedChan1()
	unbufferedChan2()
	bufferedChan()
	overtime()
}

func sendReceive() {
	// init channels
	ch1 := make(chan int)                 // 创建一个整型类型的通道
	ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
	type Equip struct{ /* 一些字段 */ }
	ch3 := make(chan *Equip)             // 创建 Equip 指针类型的通道, 可以存放 *Equip

	// 开启并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// send msg
		// 通过通道通知 main 的 goroutine
		ch1 <- 0
		ch2 <- "hello"
		eq := Equip{}
		ch3 <- &eq
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	// receive msg
	// 阻塞接收数据，执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量
	data1 := <-ch1

	// 非阻塞接收数据，使用非阻塞方式从通道接收数据时，语句不会发生阻塞
	// 非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行
	data2, ok := <-ch2

	// 接收任意数据，忽略接收的数据
	// 执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步
	<-ch3
	fmt.Println("all done")
	fmt.Println(data1, data2, ok)
}

func loopReceive() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		// 从 3 循环到 0
		for i := 3; i >= 0; i-- {
			// 发送 3 到 0 之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Print(data, " ")
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			fmt.Println()
			break
		}
	}
}

/*
  我们在将一个 channel 变量传递到一个函数时，可以通过将其指定为单向 channel 变量，从而限制该函数中可以对此 channel 的操作，比如只能往这个 channel 中写入数据，或者只能从这个 channel 读取数据
  var 通道实例 chan<- 元素类型    // 只能写入数据的通道
  var 通道实例 <-chan 元素类型    // 只能读取数据的通道
 */
func unidirectionalChan() {
	ch := make(chan int)
	// 声明一个只能写入数据的通道类型, 并赋值为 ch
	var chSendOnly chan<- int = ch
	//声明一个只能读取数据的通道类型, 并赋值为 ch
	var chRecvOnly <-chan int = ch

	go func() {
		chSendOnly <- 5
	}()

	// 判断 channel 是否已经被关闭：看第二个 bool 返回值即可，如果返回值是 false 则表示 ch 已经被关闭
	data, ok := <- chRecvOnly
	fmt.Println(data, ok)

	// 关闭 channel
	close(ch)
}

/*
  Go 语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作
 */
func unbufferedChan1() {
	// 创建一个无缓冲的通道
	court := make(chan int)
	// 计数加 2，表示要等待两个 goroutine
	wait.Add(2)
	// 启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)
	// 发球
	court <- 1
	// 等待游戏结束
	wait.Wait()
}

// wait 用来等待程序结束
var wait sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wait.Done()
	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}
		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}

func unbufferedChan2() {
	// 创建一个无缓冲的通道
	baton := make(chan int)
	// 为最后一位跑步者将计数加 1
	wait.Add(1)
	// 第一位跑步者持有接力棒
	go Runner(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wait.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒
	runner := <-baton
	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)
	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)
	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wait.Done()
		return
	}
	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)
	baton <- newRunner
}

/*
  Go 语言中有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收
  通道会阻塞发送和接收动作的条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞

  通道实例 := make(chan 通道类型, 缓冲大小)
 */
func bufferedChan() {
	// 创建一个 3 个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 发送3个整型元素到通道
	ch <- 1
	ch <- 2
	ch <- 3
	// 查看当前通道的大小
	fmt.Println(len(ch))
}

/*
  Go 语言可以使用 select 来设置超时
  虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
  与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作

  select {
    case <-chan1:
    // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
    // 如果成功向chan2写入数据，则进行该case处理语句
    default:
    // 如果上面都没有成功，则进入default处理流程
  }
  如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：
    如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复
    如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去
 */
func overtime() {
	ch := make(chan int)
	quit := make(chan bool)

	// 新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<- quit
	fmt.Println("程序结束")
}