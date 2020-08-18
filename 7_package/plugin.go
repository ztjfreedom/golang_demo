package main

import (
	"log"
	"github.com/dullgiulio/pingo"
)

func main() {
	// 从创建的可执行文件中创建一个新插件。通过 TCP 连接到它
	p := pingo.NewPlugin("tcp", "7_package/plugins/hello/hello")
	// 启动插件
	p.Start()
	// 使用完插件后停止它
	defer p.Stop()

	var resp string

	// 从先前创建的对象调用函数
	if err := p.Call("MyPlugin.SayHello", "Go developer", &resp); err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}