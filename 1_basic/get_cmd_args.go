package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
// 3个参数分别为：参数名称，默认值，-help时的说明
var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
}