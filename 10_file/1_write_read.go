package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golearn/v2/10_file/model"
	"io"
	"os"
)

func main() {
	writeJSON()
	readJSON()
	writeXML()
	readXML()
	writeGob()
	readGob()
	writeBinary()
	readBinary()
	writeTxt()
	readTxt()
}

func writeJSON() {
	info := []model.Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	// 创建文件
	filePtr, err := os.Create("10_file/data/1_info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建 Json 编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func readJSON() {
	filePtr, err := os.Open("10_file/data/1_info.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()
	var info []model.Website
	// 创建 Json 解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

func writeXML() {
	// 实例化对象
	info := model.Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}
	f, err := os.Create("10_file/data/1_info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer f.Close()
	// 序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}

func readXML() {
	// 打开 xml 文件
	file, err := os.Open("10_file/data/1_info.xml")
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return
	}
	defer file.Close()
	info := model.Website{}
	// 创建 xml 解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

/*
  为了让某个数据结构能够在网络上传输或能够保存至文件，它必须被编码然后再解码。当然已经有许多可用的编码方式了，比如 JSON、XML、Google 的 protocol buffers 等等。而现在又多了一种，由 Go 语言 encoding/gob 包提供的方式
  Gob 是 Go 语言自己以二进制形式序列化和反序列化程序数据的格式，可以在 encoding 包中找到。这种格式的数据简称为 Gob（即 Go binary 的缩写）。类似于 Python 的 “pickle” 和 Java 的 “Serialization”
 */
func writeGob() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}
	name := "10_file/data/1_info.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}

func readGob() {
	var M map[string]string
	File, _ := os.Open("10_file/data/1_info.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}

type FourByte struct {
	Num int32
}

// 虽然 Go 语言的 encoding/gob 包非常易用，而且使用时所需代码量也非常少，但是我们仍有可能需要创建自定义的二进制格式
func writeBinary(){
	file, err := os.Create("10_file/data/1_info.bin")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()
	for i := 1; i <= 3; i++ {
		info := FourByte{int32(i)}
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}

func readBinary() {
	file, err := os.Open("10_file/data/1_info.bin")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	m := FourByte{}
	for i := 1; i <= 3; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

func writeTxt() {
	// 创建一个新文件，写入内容
	filePath := "10_file/data/1_info.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	// 及时关闭
	defer file.Close()
	// 写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt 文件要看到换行效果要用 \r\n
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	// 因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	// 所以要调用 flush 方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}

func readTxt() {
	// 打开文件
	file, err := os.Open("10_file/data/1_info.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	// 及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	// 创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')  // 读到一个换行就结束
		if err == io.EOF {                         // io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}