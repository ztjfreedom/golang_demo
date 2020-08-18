package main

/*
  在 Go 语言中，不仅结构体与结构体之间可以嵌套，接口与接口间也可以通过嵌套创造出新的接口
  一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用
 */
func main() {
	// 声明写入关闭器, 并赋予 device 的实例
	var wc WriteCloser = new(device)

	// 写入数据，WriteCloser 接口嵌套了 Writer 接口，Writer 接口里定义了 Write 方法，绑定了 device 的实现
	wc.Write(nil)

	// 关闭设备
	wc.Close()

	// 声明写入器, 并赋予device的新实例
	var writeOnly Writer = new(device)

	// 写入数据
	writeOnly.Write(nil)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// 接口的嵌套
type WriteCloser interface {
	Writer
	Closer
}

// 声明一个设备结构
type device struct {
}

// 实现 Writer 的 Write() 方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 实现 Closer 的 Close() 方法
func (d *device) Close() error {
	return nil
}