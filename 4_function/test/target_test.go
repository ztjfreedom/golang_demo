package test

/*
  Go 语言自带了 testing 测试包，可以进行自动化的单元测试，输出结果验证，并且可以测试性能

  编写测试用例有以下几点需要注意：
    测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中
    测试用例的文件名必须以 _test.go 结尾
    需要使用 import 导入 testing 包
    测试函数的名称要以 Test 或 Benchmark 开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数
    单元测试则以 (t *testing.T) 作为参数，性能测试以 (t *testing.B) 做为参数
    测试用例文件使用 go test 命令来执行，源码中不需要 main() 函数作为入口，所有以 _test.go 结尾的源码文件内以 Test 开头的函数都会自动执行
 */
import "testing"

// 单元（功能）测试  go test -v
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

// 性能（压力）测试  go test -bench="."
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}

/*
  覆盖率测试  go test -cover
  覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码）
 */