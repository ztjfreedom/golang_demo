package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	read()
	write()
}

// 操作 Reader 对象的方法共有 11 个，分别是 Read()、ReadByte()、ReadBytes()、ReadLine()、ReadRune ()、ReadSlice()、ReadString()、UnreadByte()、UnreadRune()、Buffered()、Peek()
func read() {
	// Read() 方法的功能是读取数据，并存放到字节切片 p 中。Read() 执行结束会返回已读取的字节数，因为最多只调用底层的 io.Reader 一次，所以返回的 n 可能小于 len(p)
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

	// ReadByte() 方法的功能是读取并返回一个字节，如果没有字节可读，则返回错误信息
	data = []byte("C语言中文网")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)

	// ReadBytes() 方法的功能是读取数据直到遇到第一个分隔符“delim”，并返回读取的字节序列（包括“delim”）
	data = []byte("C语言, Go语言入门教程")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)

	// ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者 ReadString('\n')。ReadLine 返回一行，不包括结尾的回车字符，如果一行太长（超过缓冲区长度），参数 isPrefix 会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回
	data = []byte("Golang is a beautiful language. \r\n I like it!")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)

	// ReadRune() 方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数
	data = []byte("C语言中文网")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)

	// ReadSlice() 方法的功能是读取数据直到分隔符“delim”处，并返回读取数据的字节切片，下次读取数据时返回的切片会失效
	data = []byte("C语言中文网, Go语言入门教程")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	delim = ','
	for {
		line, err = r.ReadSlice(delim)
		fmt.Println(strings.Trim(string(line), " "))
		if err == io.EOF {
			break
		}
	}

	// ReadString() 方法的功能是读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串
	data = []byte("C语言中文网, Go语言入门教程")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	delim = ','
	str, err := r.ReadString(delim)
	fmt.Println(str, err)

	// UnreadByte() 方法的功能是取消已读取的最后一个字节（即把字节重新放回读取缓冲区的前部）。只有最近一次读取的单个字节才能取消读取
	// UnreadRune() 方法的功能是取消读取最后一次读取的 Unicode 字符。如果最后一次读取操作不是 ReadRune，UnreadRune 会返回一个错误

	// Buffered() 方法的功能是返回可从缓冲区读出数据的字节数
	data = []byte("Go语言入门教程")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	var smallBuf [14]byte
	n, err = r.Read(smallBuf[:])
	fmt.Println(string(smallBuf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println(rn)
	n, err = r.Read(smallBuf[:])
	fmt.Println(string(smallBuf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println(rn)

	// Peek() 方法的功能是读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效
	data = []byte("Go语言入门教程")
	rd = bytes.NewReader(data)
	r = bufio.NewReader(rd)
	bl, err := r.Peek(8)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(14)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
}

// 操作 Writer 对象的方法共有 7 个，分别是 Available()、Buffered()、Flush()、Write()、WriteByte()、WriteRune() 和 WriteString() 方法
func write() {
	// Available() 方法的功能是返回缓冲区中未使用的字节数
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())

	// Buffered() 方法的功能是返回已写入当前缓冲区中的字节数
	// Flush() 方法的功能是把缓冲区中的数据写入底层的 io.Writer，并返回错误信息
	wr = bytes.NewBuffer(nil)
	w = bufio.NewWriter(wr)
	p = []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
	w.Flush()
	fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())

	// Write() 方法的功能是把字节切片 p 写入缓冲区，返回已写入的字节数
	wr = bytes.NewBuffer(nil)
	w = bufio.NewWriter(wr)
	p = []byte("C语言中文网")
	n, err := w.Write(p)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)

	// WriteByte() 方法的功能是写入一个字节，如果成功写入，error 返回 nil，否则 error 返回错误原因
	wr = bytes.NewBuffer(nil)
	w = bufio.NewWriter(wr)
	var c byte = 'G'
	err = w.WriteByte(c)
	w.Flush()
	fmt.Println(string(wr.Bytes()), err)

	// WriteRune() 方法的功能是以 UTF-8 编码写入一个 Unicode 字符，返回写入的字节数和错误信息
	wr = bytes.NewBuffer(nil)
	w = bufio.NewWriter(wr)
	var r rune = 'G'
	size, err := w.WriteRune(r)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)

	// WriteString() 方法的功能是写入一个字符串，并返回写入的字节数和错误信息
	wr = bytes.NewBuffer(nil)
	w = bufio.NewWriter(wr)
	s := "C语言中文网"
	n, err = w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}