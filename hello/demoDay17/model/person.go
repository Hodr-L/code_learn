package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func Newperson(name string) *person {
	return &person{
		Name: name,
	}
}
func (p *person) Setage_sal(age int, sal float64) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入年龄错误！")
		return
	}
	if sal < 0 {
		fmt.Println("输入工资错误！")
	} else {
		p.sal = sal
	}

}
func (p *person) Getage() int {
	return p.age
}
func (p *person) Getsal() float64 {
	return p.sal
}
func (p *person) Getname() string {
	return p.Name
}
