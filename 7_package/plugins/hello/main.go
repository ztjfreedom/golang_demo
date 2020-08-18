// 创建新的二进制文件
package main

import "github.com/dullgiulio/pingo"

// 创建要导出的对象
type MyPlugin struct{}

// 导出的方法，带有 rpc 签名
func (p *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "Hello, " + name
	return nil
}

/*
  Pingo 是一个用来为Go语言程序编写插件的简单独立库，因为 Go 本身是静态链接的，因此所有插件都以外部进程方式存在。Pingo 旨在简化标准 RPC 包，支持 TCP 和 Unix 套接字作为通讯协议
  go build
 */
func main() {
	plugin := &MyPlugin{}

	// 注册要导出的对象
	pingo.Register(plugin)
	// 运行主程序
	pingo.Run()
}