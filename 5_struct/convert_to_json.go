package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	structToJson()
}

func structToJson() {
	f := Fighter {
		Name: "cow boy",
		Age:  37,
		Skills: []Skill {
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	result, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}
	jsonStringData := string(result)
	fmt.Println(jsonStringData)
}

type Skill struct {
	Name  string `json:"SkillName"`  // 在转换 JSON 格式时，JSON 的各个字段名称默认使用结构体的名称，如果想要指定为其它的名称使用标签
	Level int
}

type Fighter struct {
	Name   string
	Age    int
	Skills []Skill
}