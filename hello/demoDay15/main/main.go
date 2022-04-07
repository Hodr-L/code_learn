package main

import (
	"fmt"
	"hello/demoDay15/testFunc"
)

func main() {
	//var p testFunc.Num
	//fmt.Println("请输入：")
	//fmt.Scanln(&p.Name)
	//p.Jisuan(b, c)

	//stu := testFunc.Circle{
	//	Name: "tom",
	//	Age:  20,
	//}
	//fmt.Println(&stu)

	var a testFunc.MethodUtils
	var m, n int
	fmt.Println("请输入m和n：")
	fmt.Scan(&m, &n)
	l := a.Test(m, n)
	fmt.Println(l)
}
