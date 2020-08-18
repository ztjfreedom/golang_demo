package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
)

/*
  使用 Go 语言开发一些程序的时候，往往出现多个进程同时操作同一份文件的情况，这很容易导致文件中的数据混乱。这时我们就需要采用一些手段来平衡这些冲突，文件锁（flock）应运而生
  对于 flock，最常见的例子就是 Nginx，进程运行起来后就会把当前的 PID 写入这个文件，当然如果这个文件已经存在了，也就是前一个进程还没有退出，那么 Nginx 就不会重新启动，所以 flock 还可以用来检测进程是否存在
  flock 是对于整个文件的建议性锁。也就是说，如果一个进程在一个文件（inode）上放了锁，其它进程是可以知道的（建议性锁不强求进程遵守）。最棒的一点是，它的第一个参数是文件描述符，在此文件描述符关闭时，锁会自动释放。而当进程终止时，所有的文件描述符均会被关闭。所以很多时候就不用考虑类似原子锁解锁的事情

  flock 属于建议性锁，不具备强制性。一个进程使用 flock 将文件锁住，另一个进程可以直接操作正在被锁的文件，修改文件中的数据，原因在于 flock 只是用于检测文件是否被加锁，针对文件已经被加锁，另一个进程写入数据的情况，内核不会阻止这个进程的写入操作，也就是建议性锁的内核处理策略
  flock 主要三种操作类型：
    LOCK_SH：共享锁，多个进程可以使用同一把锁，常被用作读共享锁
    LOCK_EX：排他锁，同时只允许一个进程使用，常被用作写锁
    LOCK_UN：释放锁

  下面的程序需要在 Linux 或 Mac 系统下才能正常运行
 */
func main() {
	test_file_path, _ := os.Getwd()  // Getwd returns a rooted path name corresponding to the current directory
	locked_file := test_file_path

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			flock := New(locked_file)
			err := flock.Lock()
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("output : %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}

// 文件锁
type FileLock struct {
	dir string
	f   *os.File
}

func New(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}

// 加锁
// 进程使用 flock 尝试锁文件时，如果文件已经被其他进程锁住，进程会被阻塞直到锁被释放掉
func (l *FileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

// 释放锁
// flock 锁的释放非常具有特色，即可调用 LOCK_UN 参数来释放文件锁，也可以通过关闭 fd 的方式来释放文件锁（flock 的第一个参数是 fd），意味着 flock 会随着进程的关闭而被自动释放掉
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}