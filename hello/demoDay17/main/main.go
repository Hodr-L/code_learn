package main

import (
	"fmt"
	"hello/demoDay17/model"
)

func main() {
	//bank.Operate()

	//var a string
	//var b int
	//var c float64
	//stu := model.Newperson(a)
	//for {
	//
	//	fmt.Println("输入姓名：")
	//	fmt.Scan(&stu.Name)
	//	fmt.Println("输入年龄：")
	//	fmt.Scan(&b)
	//	fmt.Println("输入工资：")
	//	fmt.Scan(&c)
	//	stu.Setage_sal(b, c)
	//	if stu.Getage() > 0 && stu.Getage() < 150 {
	//		if stu.Getsal() > 0 {
	//			fmt.Println()
	//			fmt.Println(stu.Getname())
	//			fmt.Println(stu.Getage())
	//			fmt.Println(stu.Getsal())
	//		}
	//	}
	//}

	var a string
	var b string
	var c float64
	user := model.User(a, b, c)
	fmt.Println("请输入用户名：")
	fmt.Scan(&a)
	if len(a) > 5 && len(a) < 11 {
		user.Setname(a)
	} else {
		fmt.Println("用户名长度不正确！")
		return
	}
	fmt.Println("请输入密码：")
	fmt.Scan(&b)
	if len(b) != 6 {
		fmt.Println("密码长度不正确！")
		return
	} else {
		user.Setpwd(b)
	}
	fmt.Println("请输入余额：")
	fmt.Scan(&c)
	if c < 20 {
		fmt.Println("余额大小不正确！")
		return
	} else {
		user.Setsal(c)
	}
	fmt.Printf("用户名：%v\n密码:%v\n余额：%v\n", user.Getname(), user.Getpwd(), user.Getsal())
}
