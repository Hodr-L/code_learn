package main

import (
	"fmt"
)

func input() {
	var name string
	var age int
	var scorce float32
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：：")
	fmt.Scanln(&age)
	fmt.Println("请输入分数：：")
	fmt.Scanln(&scorce)

	if scorce >= 60.0 {
		fmt.Println("\n姓名：", name, "\n年龄：", age, "\n分数：", scorce, "\n考试及格")
	} else {
		fmt.Println("姓名：", name, "\n年龄：", age, "\n分数：", scorce, "\n考试不及格")
	}

}
func def() {
	defer fmt.Println("1")
	fmt.Println("2")
}
func main() {
	def()
}
