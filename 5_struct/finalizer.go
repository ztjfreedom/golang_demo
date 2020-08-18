package main

import (
	"log"
	"runtime"
	"time"
)

/*
  Go 语言自带垃圾回收机制（GC）。GC 通过独立的进程执行，它会搜索不再使用的变量，并将其释放。需要注意的是，GC 在运行时会占用机器资源
  GC 是自动进行的，如果要手动进行 GC，可以使用 runtime.GC() 函数，显式的执行 GC。显式的进行 GC 只在某些特殊的情况下才有用，比如当内存资源不足时调用 runtime.GC() ，这样会立即释放一大片内存，但是会造成程序短时间的性能下降
  finalizer（终止器）是与对象关联的一个函数，通过 runtime.SetFinalizer 来设置，如果某个对象定义了 finalizer，当它被 GC 时候，这个 finalizer 就会被调用，以完成一些特定的任务，例如发信号或者写日志等
  也可以使用 SetFinalizer(x, nil) 来清理绑定到 x 上的终止器
 */
func main() {

	entry()

	// 运行一次后对象 r 就被内存回收了，所以 findRoad 只会被调用一次
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}

}

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)  // 与对象 r 关联，绑定了 findRoad 函数，SetFinalizer 函数可以将 x 的终止器设置为 f，当垃圾收集器发现 x 不能再直接或间接访问时，它会清理 x 并调用 f(x)
}