package main

import (
	"fmt"
	"time"
)

/*
  UTC + 时区差 ＝ 本地时间
  在Go语言的 time 包里面有两个时区变量，如下：
    time.UTC：UTC 时间
    time.Local：本地时间
 */
func main() {
	now()
	timeStamp()
	timeStampToTime()
	weekday()
	add()
	format()
	parseStringToTime()
	timer()
}

func now() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 时
	minute := now.Minute() // 分
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timeStamp() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)
}

func timeStampToTime() {
	now := time.Now()                  // 获取当前时间
	timestamp := now.Unix()            // 时间戳
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 时
	minute := timeObj.Minute() // 分
	second := timeObj.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func weekday() {
	t := time.Now()
	fmt.Println(t.Weekday().String())
}

// 还有 Sub, Equal, Before, After 等函数
func add() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加 1 小时后的时间
	fmt.Println(later)
}

func format() {
	now := time.Now()
	// 格式化的模板为 Go 的出生时间 2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func parseStringToTime() {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)
}

func timer() {
	ticker := time.Tick(1 * time.Second) //定义一个 1 秒间隔的定时器
	count := 1
	for i := range ticker {
		fmt.Println(i, count)  // 每秒都会执行的任务
		if count == 3 {
			return
		}
		count ++
	}
}