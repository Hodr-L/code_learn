package main

import (
	"fmt"
)

func control(a int, b float64, e bool) {
	if a >= 180 && b >= 1 && e == true {
		fmt.Println("嫁！")
	} else if (a >= 180 && b >= 1) || (b >= 1 && e == true) || (a >= 180 && e == true) {
		fmt.Println("嫁吧！")
	} else {
		fmt.Println("滚！")
	}
}
func main() {
	var a int
	var b float64
	var c bool
	fmt.Println("请输入身高（cm）：")
	_, _ = fmt.Scanln(&a)
	fmt.Println("请输入资产（千万）：")
	_, _ = fmt.Scanln(&b)
	fmt.Println("是否帅（true/false）：")
	_, _ = fmt.Scanln(&c)
	d := a
	e := b
	f := c
	control(d, e, f)
}
