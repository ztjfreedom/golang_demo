package model

import "fmt"

// 首字母小写其它包不能直接访问
type person struct {
	Name string
	age int
	sal float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}
// 为了访问 age 和 sal 我们编写一对 SetXxx 的方法和 GetXxx 的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 1000 && sal <= 300000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}