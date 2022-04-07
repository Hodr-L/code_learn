package main

import (
	"fmt"
	"unicode"
)

// run command
// go run ./src/chapter4/work4_6/work4_6.go
func main() {
	rs := []rune{'H', 'e', 'l', 'l', 'o', ' ', ' ', ' ', '世', '界'}  //定义一个rune类型数组，并填充初始值
	fmt.Println("input string:\t\n", string(rs))                    //打印数组内容
	bs := []byte(string(rs))                                        //定义byte类型数组，将rs数组存入，转为UTF-8编码
	fmt.Println("output string:\t\n", string(uniqueSpaceSlice(bs))) //打印参数为bs的uniqueSpaceSlice函数的结果
}

func uniqueSpaceSlice(bs []byte) []byte { //声明函数，且参数为byte类型数组 bs
	fmt.Println("input []bytes:\t\n", bs) //打印rs的UTF-8编码，即为byte类型的bs
	flag := 0                             // 连续出现space次数的flag
	for i := 0; i < len(bs); i++ {
		if unicode.IsSpace(rune(bs[i])) {
			flag++
			if flag > 1 {
				bs = append(bs[:i], bs[i+1:]...)
				i-- // we delete a element in slice, so we should make i minus 1
			}
		} else {
			flag = 0
		}
	}
	fmt.Println("output []bytes:\t", bs)
	return bs
}
