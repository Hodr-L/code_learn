package main

import (
	"fmt"
	"hello/demoDay16/face"
)

func main() {
	//var student = face.Student{
	//	Name:   "杨鑫",
	//	Gender: "女",
	//	Age:    18,
	//	Id:     18001040306,
	//	Score:  96.3,
	//}
	//fmt.Println(student.Say())

	//var dog = face.Dog{
	//	Name:   "小花",
	//	Age:    2,
	//	Weid: 12.6,
	//}
	//fmt.Println(dog.Say())

	//	var box face.Box
	//	var x, y, z float64
	//	fmt.Println("请输入长、宽、高：")
	//	fmt.Scan(&x, &y, &z)
	//	box.Len = x
	//	box.Wide = y
	//	box.High = z
	//	fmt.Println(box.Say())

	//for {
	//	var a string
	//	var b int
	//	fmt.Println("请输入姓名：")
	//	fmt.Scan(&a)
	//	if a == "n" {
	//		fmt.Println("退出程序")
	//		break
	//	}
	//	fmt.Println("请输入年龄：")
	//	fmt.Scan(&b)
	//	var stu = face.Newjingqu(a, b)
	//	fmt.Println(stu.Say)
	//}
	var a string
	var b int
	var stu = face.Stu(a, b)
	fmt.Println("输入name：")
	fmt.Scan(&a)
	fmt.Println("输入age：")
	fmt.Scan(&b)
	fmt.Println()
	stu.Setname(a)
	stu.Setage(b)

	fmt.Println(stu.Getname())
	fmt.Println(stu.Getage())
}
