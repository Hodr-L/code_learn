package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point float64 = 10.1
	a = point
	//将a付给一个point变量
	if b, ok := a.(float64); ok {
		fmt.Println("ok")
		fmt.Println(b)
	} else {
		fmt.Println("no")
	}
	fmt.Println("继续执行...")
}
