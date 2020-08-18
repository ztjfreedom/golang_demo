package main

import "fmt"

func main() {
	condition()
	loop()
	forRange()
	switchCase()
	goTo()
	breakToLabel()
	continueToLabel()
}

func condition() {
	if err := connect(); err != nil {
		fmt.Println(err)
	}
}

func connect() interface{} {
	return "Connection Error"
}

func loop() {
	// Go 语言中的循环语句只支持 for 关键字，而不支持 while 和 do-while 结构
	var count int
	for count <= 10 {
		count ++
	}
	fmt.Println(count)

	sum := 0
	for i := 0; i < 10; i++ {  // for 后面的条件表达式不需要用圆括号括起来
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	for {  // 不需要 while (true) 之类的来无限循环
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

	ILoop:
		for i := 0; i < 5; i++ {
			for j := 0; j < 10; j++ {
				if j > 5 {
					break ILoop  // 可以指定跳出到哪层循环
				}
			}
			fmt.Println("Break J")
		}
	fmt.Println("Break I")
}

func forRange() {
	/*
	  for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句
	  for key, val := range coll {
		  ...
	  }
	  需要要注意的是，val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。一个字符串是 Unicode 编码的字符（或称之为 rune）集合，因此也可以用它来迭代字符串

	  通过 for range 遍历的返回值有一定的规律：
	    数组、切片、字符串返回索引和值
	    map 返回键和值
	    通道（channel）只返回通道内的值
	 */
	// slice
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d  ", key, value)
	}
	fmt.Println()

	// string
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x %s  ", key, value, string(value))  // 代码中的变量 value，实际类型是 rune 类型，以十六进制打印出来就是字符的编码
	}
	fmt.Println()

	// map
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Printf("key:%s value:%d  ", key, value)
	}
	fmt.Println()

	for _, value := range m {
		fmt.Printf("value:%d  ", value)
	}
	fmt.Println()

	// channel
	// for range 可以遍历通道（channel），但是通道在遍历时，只输出一个值，即管道内的类型对应的数据
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func switchCase() {
	// Go 语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	case "bro":
		fmt.Println("brother")
	default:
		fmt.Println("who")
	}

	var r = 12
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	// Go 语言中 case 是一个独立的代码块，执行完毕后不会像 C 语言那样紧接着执行下一个 case，但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能
	// 新编写的代码，不建议使用 fallthrough
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

func goTo() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto BreakHere
			}
		}
		fmt.Println("If not break")
	}
	// 手动返回, 避免执行进入标签，此处如果不手动返回，在不满足条件时，也会执行标签中的代码
	return
BreakHere:
	fmt.Println("Done Break Label")

	// 使用 goto 集中处理错误
	err := firstCheckError()
	if err != nil {
		goto onExit
	}
	err = secondCheckError()
	if err != nil {
		goto onExit
	}
	fmt.Println("done")
	return
onExit:
	fmt.Println(err)
}

func firstCheckError() interface{} {
	return nil
}

func secondCheckError() interface{} {
	return "Error"
}

func breakToLabel() {
	// Go 语言中 break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Print(i, j, "  ")
				break OuterLoop
			case 3:
				fmt.Print(i, j, "  ")
				break OuterLoop
			}
		}
	}
	fmt.Println()
}

func continueToLabel() {
	// Go 语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Print(i, j, "  ")
				continue OuterLoop
			}
		}
	}
	fmt.Println()
}