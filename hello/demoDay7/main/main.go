package main

import (
	"fmt"
	"hello/demoDay7/sum"
)

func main() {
	////fmt.Println(sum.Sum(10, 10, 10, 10))
	//f := sum.Addupper()
	//fmt.Println(f(1))
	//f2 := sum.Makesuffix(".jpg")
	//fmt.Println("文件后缀处理=", f2("winter"))

	var a int
	fmt.Println("请输入：")
	fmt.Scan(&a)
	sum.Sum4(a)
}
