package main

import "fmt"

func enter(a int, b int, c int) {
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	if c > max {
		max = c
	}
	fmt.Println("最大值为：", max)
}

func main() {
	var c int
	var d int
	var e int
	fmt.Println("请输入a、b、c的值：")
	fmt.Scanln(&c, &d, &e)
	enter(c, d, e)
}
