package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
)

func main() {
	execute()
	currentUser()
	signalProcess()
}

/*
  常用函数：Hostname Environ Getenv Setenv Exit Getuid Getgid Getpid Getwd Mkdir MkdirAll Remove
 */

// exec 包可以执行外部命令
func execute() {
	cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func currentUser() {
	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	// 用户所在的所有的组的 id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)
}

// 一个运行良好的程序在退出（正常退出或者强制退出，如 Ctrl+C，kill 等）时是可以执行一段清理代码的，将收尾工作做完后再真正退出。一般采用系统 Signal 来通知系统退出
// 运行该程序，然后在 CMD 窗口中通过 Ctrl+C 来结束该程序，便会得到输出结果
func signalProcess() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}