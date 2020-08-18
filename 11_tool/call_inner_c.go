package main

/*
#include <stdio.h>
#include <stdlib.h>
typedef struct {
    int id;
}ctx;
ctx *createCtx(int id) {
    ctx *obj = (ctx *)malloc(sizeof(ctx));
    obj->id = id;
    return obj;
}
*/
import "C"

import (
	"fmt"
)

/*
  在 Go 语言的源代码中直接声明 C 语言代码是比较简单的应用情况，可以直接使用这种方法将 C 语言代码直接写在 Go 语言代码的注释中，并在注释之后紧跟 import "C"，通过 C.xx 来引用 C 语言的结构和函数

  也可以通过封装实现 C++ 接口的调用
  C++ 代码需要提前编译成动态库（拷贝到系统库目录可以防止 go 找不到动态库路径），go 程序运行时会去链接
 */
func main() {
	var ctx *C.ctx = C.createCtx(100)
	fmt.Printf("id : %d\n", ctx.id)
}