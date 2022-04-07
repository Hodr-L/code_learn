package main

import "fmt"

func reverse(a string) {
	//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	//		//初始i，j为0，s数组长度减1(防止数组越界)，若i<j,则i增加1，j减少1
	//		s[i], s[j] = s[j], s[i]
	//		//s[i]=s[j],s[j]=s[i]
	//	}
	//
}

func main() {
	//a := [...]int{1, 2, 3, 4, 5, 6, 7}
	//reverse(a[0:])
	//fmt.Println(a)
	var a string
	fmt.Println("输入：")
	fmt.Scan(&a)
	reverse(a)
}
