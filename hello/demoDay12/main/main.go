package main

import (
	"fmt"
	"hello/demoDay12/chazhao"
)

func main() {
	//var heroName = ""
	//fmt.Println("请输入：")
	//fmt.Scan(&heroName)
	//chazhao.Shunxu2(heroName)
	arr := []int{1, 2, 3, 4, 5, 6}
	var a int
	fmt.Println("请输入：")
	fmt.Scanf("%d", &a)
	chazhao.ErfenTest(&arr, 0, len(arr)-1, a)
}
