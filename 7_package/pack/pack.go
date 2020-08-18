package pack

import F "fmt"

func PackPrint(data interface{}) {
	F.Println(data)
}

/*
  init() 函数的特性如下：
    每个源码可以使用 1 个 init() 函数
    init() 函数会在程序执行前（main() 函数执行前）被自动调用
    调用顺序为 main() 中引用的包，以深度优先顺序初始化
    同一个包中的多个 init() 函数的调用顺序不可预期
    init() 函数不能被其他函数调用

  假设有这样的包引用关系：main → A → B → C，那么这些包的 init() 函数调用顺序为：
  C.init → B.init → A.init → main
 */
func init() {
	F.Println("Pack Init")
}