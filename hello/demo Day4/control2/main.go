package main

import "fmt"

func control(a int, b int) {
	if a >= 4 && a <= 10 {
		if b >= 18 && b <= 60 {
			fmt.Println("旺季：\n成人票：60元")
		} else if b > 60 {
			fmt.Println("旺季：\n 老人票：20元")
		} else {
			fmt.Println("旺季：\n儿童票：30元")
		}
	} else {
		if b >= 18 && b <= 60 {
			fmt.Println("淡季：\n成人票：40元")
		} else {
			fmt.Println("淡季：\n其他：20元")
		}
	}
}
func main() {
	var a int
	var b int
	fmt.Println("请输入月份：")
	_, _ = fmt.Scanln(&a)
	fmt.Println("请输入年龄：")
	_, _ = fmt.Scanln(&b)
	control(a, b)
}
