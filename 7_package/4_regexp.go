package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	reg1()
	reg2()
	reg3()
	reg4()
	reg5()
	regReplace()
}

func reg1() {
	buf := "abc azc a7c aac 888 a9c  tac"
	// 解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	// 根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func reg2() {
	buf := "abc azc a7c aac 888 a9c  tac"

	reg1 := regexp.MustCompile(`a[0-9]c`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func reg3() {
	buf := "abc azc a7c aac 888 a9c  tac"

	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func reg4() {
	buf := "43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8   "

	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
}

func reg5() {
	// 原生字符串
	buf := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>C语言中文网 | Go语言入门教程</title>
</head>
<body>
    <div>Go语言简介</div>
    <div>Go语言基本语法
    Go语言变量的声明
    Go语言教程简明版</div>
    <div>Go语言容器</div>
    <div>Go语言函数</div>
</body>
</html>
    `

	//reg := regexp.MustCompile(`<div>.*</div>`)
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)  // 分组，并且可以匹配到多行
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	result := reg.FindAllStringSubmatch(buf, -1)

	for _, text := range result {
		fmt.Println("text[1] = ", text[1])
	}
}

func regReplace() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"          // 正则

	f := func(s string) string{
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v * 2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}