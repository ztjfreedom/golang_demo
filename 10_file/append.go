package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
  O_RDONLY：只读模式打开文件
  O_WRONLY：只写模式打开文件
  O_RDWR：读写模式打开文件
  O_APPEND：写操作时将数据附加到文件尾部（追加）
  O_CREATE：如果不存在将创建一个新文件
  O_EXCL：和 O_CREATE 配合使用，文件必须不存在，否则返回一个错误
  O_SYNC：当进行一系列写操作时，每次都要等待上次的 I/O 操作完成再进行
  O_TRUNC：如果可能，在打开时清空文件
 */
func main() {
	createFile()
	appendFile()
	readAppend()
	copyFile()
}

// 创建一个新文件，写入内容“http://c.biancheng.net/golang/”
func createFile() {
	filePath := "10_file/data/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 及时关闭 file 句柄
	defer file.Close()
	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString("http://c.biancheng.net/golang/ \n")
	// Flush 将缓存的文件真正写入到文件中
	write.Flush()
}

// 打开一个存在的文件，在原来的内容追加内容“C语言中文网”
func appendFile() {
	filePath := "10_file/data/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)  // append
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("C语言中文网 \r\n")
	write.Flush()
}

// 打开一个存在的文件，将原来的内容读出来，显示在终端，并且追加“Hello，C语言中文网”
func readAppend() {
	filePath := "10_file/data/golang.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	// 读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	// 写入文件
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("Hello，C语言中文网。 \r\n")
	}
	write.Flush()
}

// 将一个文件的内容复制到另外一个文件
func copyFile() {
	file1Path := "10_file/data/golang.txt"
	file2Path := "10_file/data/golang_copy.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
	}
}