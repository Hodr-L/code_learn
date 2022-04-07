package slice

import "fmt"

func Slice2() {
	str := "hello@world"
	//slice := str[6:]
	//fmt.Println(slice)
	a := []rune(str)
	a[0] = '你'
	str = string(a)
	fmt.Println(str)

}
