package slice

import "fmt"

func Slice() { //make 的使用
	//var slice = make([]int, 4)
	//slice[0] = 1
	//slice[1] = 2
	//var slice = make([]int, 1)
	//slice = append(slice, 1, 2, 3, 4, 5)
	var slice = []int{1, 2, 3, 4, 5}
	a := make([]int, 1)
	copy(a, slice)
	for _, v := range a {
		fmt.Println(v)
	}
	//fmt.Printf("\n")
	//fmt.Println(cap(a))
}
