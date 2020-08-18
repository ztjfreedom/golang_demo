package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bufioWrite()
	bufioRead()
}

func bufioWrite() {
	name := "10_file/data/buffer.txt"
	content := "http://c.biancheng.net/golang/"
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer fileObj.Close()
	writeObj := bufio.NewWriterSize(fileObj, 4096)
	// 使用 Write 方法，需要使用 Writer 对象的 Flush 方法将 buffer 中的数据刷到磁盘
	buf := []byte(content)
	if _, err := writeObj.Write(buf); err == nil {
		if err := writeObj.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("数据写入成功")
	}
}

func bufioRead() {
	fileObj, err := os.Open("10_file/data/buffer.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fileObj.Close()
	// 一个文件对象本身是实现了 io.Reader 的，使用 bufio.NewReader 去初始化一个 Reader 对象，存在 buffer 中的，读取一次就会被清空
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	// 读取 Reader 对象中的内容到 []byte 类型的 buf 中
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取的字节数:" + strconv.Itoa(info))
	// 这里的 buf 是一个 []byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("读取的文件内容:", string(buf))
}