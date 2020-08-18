package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(GetSingleInsByLock())
	fmt.Println(NewConfig())
	fmt.Println(GetSingleInsByDoubleCheck())
	fmt.Println(GetInstance())
}

// 1. 使用锁
type Tool struct {
	values int
}

var lock sync.Mutex

var instance *Tool

func GetSingleInsByLock() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

// 2. 使用常驻内存
type Cfg struct {}

var cfg *Cfg = new(Cfg) // 使用 init 或全局变量
//func init()  {
//	cfg = new(Cfg)
//}

func NewConfig() *Cfg {
	return cfg
}

// 3. 双重检查
func GetSingleInsByDoubleCheck() *Tool {
	// 第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = new(Tool)
		}
		lock.Unlock()
	}
	return instance
}

// 4. sync.Once
// sync.Once 内部本质上也是双重检查的方式，但在写法上会比自己写双重检查更简洁
var once sync.Once

func GetInstance() *Tool {
	once.Do(func() {
		instance = new(Tool)
	})
	return instance
}