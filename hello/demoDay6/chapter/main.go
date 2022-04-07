package main

import (
	"fmt"
	"hello/demoDay6/aaaa"
)

func test1(a float64, b float64, c string) float64 {
	var res float64
	switch c {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}
func main() {
	//var a float64
	//var b float64
	//var c string
	//fmt.Println("输入：")
	//fmt.Scanln(&a, &c, &b)
	//res := test1(a, b, c)
	//fmt.Println(res)
	fmt.Println(aaaa.Res(5))
}
