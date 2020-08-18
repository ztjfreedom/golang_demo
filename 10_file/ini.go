package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
  INI 文件是 Initialization File 的缩写，即初始化文件，是 Windows 的系统配置文件所采用的存储格式，统管 Windows 的各项配置。INI 文件格式由节（section）和键（key）构成，一般用于操作系统、虚幻游戏引擎、GIT 版本管理中，这种配置文件的文件扩展名为 .ini
  INI 文件由多行文本组成，整个配置由 [ ] 拆分为多个“段”（section）。每个段中又以 = 分割为“键”和“值”
  INI 文件以 ; 置于行首视为注释，注释后将不会被处理和识别
 */

func main() {
	fmt.Println(getValue("10_file/config.ini", "remote \"origin\"", "fetch"))
	fmt.Println(getValue("10_file/config.ini", "core", "hideDotFiles"))
}

// 根据文件名，段名，键名获取 ini 的值
func getValue(filename, expectSection, expectKey string) string {
	// 打开文件
	file, err := os.Open(filename)
	// 文件找不到，返回空
	if err != nil {
		return ""
	}
	// 在函数结束时，关闭文件
	defer file.Close()
	// 使用读取器读取文件
	reader := bufio.NewReader(file)
	// 当前读取的段的名字
	var sectionName string
	for {
		// 读取文件的一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 切掉行的左右两边的空白字符
		linestr = strings.TrimSpace(linestr)
		// 忽略空行
		if linestr == "" {
			continue
		}
		// 忽略注释
		if linestr[0] == ';' {
			continue
		}
		// 行首和尾巴分别是方括号的，说明是段标记的起止符
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			// 将段名取出
			sectionName = linestr[1 : len(linestr)-1]
			// 这个段是希望读取的
		} else if sectionName == expectSection {
			// 切开等号分割的键值对
			pair := strings.Split(linestr, "=")
			// 保证切开只有 1 个等号分割的简直情况
			if len(pair) == 2 {
				// 去掉键的多余空白字符
				key := strings.TrimSpace(pair[0])
				// 是期望的键
				if key == expectKey {
					// 返回去掉空白字符的值
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}