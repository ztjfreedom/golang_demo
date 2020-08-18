package main

import (
	"flag"
	"fmt"
)

/*
  几个概念：
    命令行参数（或参数）：是指运行程序时提供的参数
    已定义命令行参数：是指程序中通过 flag.Type 这种形式定义了的参数
    非 flag（non-flag）命令行参数（或保留的命令行参数）：可以简单理解为 flag 包不能解析的参数

  有以下两种常用的定义命令行 flag 参数的方法
    flag.Type()
    flag.TypeVar()
 */
var Input_pstrName = flag.String("name", "gerry", "input ur name")  // flag.Type()
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")  // flag.TypeVar()
}

func main() {
	Init()
	flag.Parse()
	// After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1
	// Args returns the non-flag command-line arguments
	// NArg is the number of arguments remaining after flags have been processed
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)
}