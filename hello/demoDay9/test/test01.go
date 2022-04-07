package test

import (
	"fmt"
	"math/rand"
	"time"
)

func Test01() {
	//var a int
	//var i = 0
	fmt.Println("猜数字！")

	//for a < 10 {
	rand.Seed(time.Now().UnixNano())
	b := rand.Intn(100) + 1
	for i := 0; i < 10; i++ {
		fmt.Println(b)
	}
	//	fmt.Printf("请输入一个数字：")
	//	fmt.Scan(&a)
	//	if a != b {
	//		fmt.Println("猜错了！")
	//		i++
	//		if 10-i != 0 {
	//			fmt.Printf("还剩%d次机会！\n", 10-i)
	//		} else {
	//			fmt.Println("没机会了，下次加油！")
	//		}
	//	} else {
	//		fmt.Println("猜对了！")
	//	}
	//}
}

//}
