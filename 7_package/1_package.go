package main

import (
	_ "database/sql" // 只是希望执行包初始化的 init 函数，而不使用包内部的数据时，可以使用匿名引用格式
	"fmt"
	"golearn/v2/7_package/pack" // go module 的形式
	"golearn/v2/7_package/model"
)

/*
  包的特性如下：
    一个目录下的同级文件归属一个包
    包名可以与其目录不同名
    包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件
 */
func main() {
	pack.PackPrint("A")
	modelInstance()
}

func modelInstance() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
}