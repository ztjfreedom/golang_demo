package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 在有 -v 标识的时候周期性的输出当前目录的总和，如果只想看到最终的结果省略 -v 即可
var verbose = flag.Bool("v", false, "显示详细进度")

// 为每一个 walkDir 的调用创建一个新的 goroutine。它使用 sync.WaitGroup 来为当前存活的 walkDir 调用计数，一个 goroutine 在计数器减为 0 的时候关闭 fileSizes 通道
// go run main.go <directory>
func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 并行遍历每一个文件树
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// 定期打印结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // 最终总数
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1024)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回 directory 目录中的条目
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}         // 获取令牌
	defer func() { <-sema }()  // 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}