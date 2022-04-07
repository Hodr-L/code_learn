package main

import (
	"fmt"
)

func test(b byte) byte {
	return b + 1
}
func control(f byte) {
	switch test(f) {
	case 'a', 'h':
		fmt.Println("星期一！")
	case 'b':
		fmt.Println("星期二！")
	case 'c':
		fmt.Println("星期三！")
	case 'd':
		fmt.Println("星期四！")
	case 'e':
		fmt.Println("星期五！")
	case 'f':
		fmt.Println("星期六！")
	default:
		fmt.Println("星期日！")
	}
}
func con(score int) {
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("优秀！")
		fallthrough
	case score >= 60 && score < 90:
		fmt.Println("良！")
	case score < 60 && score >= 0:
		fmt.Println("不及格！")
	default:
		fmt.Println("error")
	}
}
func inter(i int) {
	var x interface{}
	x = i
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x的类型是:int")
	case float64:
		fmt.Printf("x的类型是:float64")
	case func(int) float64:
		fmt.Printf("x的类型是:fuc(int)")
	case bool, string:
		fmt.Printf("x的类型是:bool或string")
	default:
		fmt.Printf("x的类型是未知类型")
	}
}

func main() {
	////var a byte
	//var a int
	//fmt.Println("请输入字符：")
	//_, _ = fmt.Scanf("%d", &a)
	////control(a)
	//con(a)
	var i int
	fmt.Printf("请输入一个数：")
	_, _ = fmt.Scanf("%d", &i)

	inter(i)
}
