package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	writeZip()
	readZip()
	writeTar()
	readTar()
}

func writeZip() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	w := zip.NewWriter(buf)

	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile("10_file/data/2_info.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
}

func readZip() {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("10_file/data/2_info.zip")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer r.Close()
	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("文件名: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Printf(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			fmt.Printf(err.Error())
		}
		rc.Close()
	}
}

func writeTar() {
	f, err := os.Create("10_file/data/2_info.tar") // 创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	fileInfo, err := os.Stat("10_file/data/1_info.txt") // 获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
	}

	err = tw.WriteHeader(hdr) // 写入头文件信息
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Open("10_file/data/1_info.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) // 将文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func readTar() {
	f, err := os.Open("10_file/data/2_info.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("10_file/data/2_info.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}