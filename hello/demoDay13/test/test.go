package test

import (
	"fmt"
	"math/rand"
	"time"
)

func Test1() {
	var b = [10]int{}
	var min int
	var max int
	var bestmin int
	var and float64 = 0
	var index = -1

	rand.Seed(time.Now().UnixNano()) //定义种子

	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(100) + 1
	} //设置随机数

	for c := 0; c < len(b)-1; c++ {
		for a := 0; a < len(b)-1; a++ {
			if b[a] < b[a+1] {
				min = b[a]
				b[a] = b[a+1]
				b[a+1] = min
				//最大和最小
				max = b[0]
				bestmin = b[len(b)-1]
			}
			if b[a] == 55 {
				index = a
			}
		}

	} //冒泡排序

	for _, c := range b {
		and += float64(c)
	} //求和

	fmt.Println("倒序输出：", b)
	fmt.Println("最大值为：", max)
	fmt.Println("最小值为：", bestmin)
	fmt.Println("平均值为：", and/10)
	//判断是否存在55
	if index == -1 {
		fmt.Println("不存在55")
	} else {
		fmt.Println("55的下标为：", index)
	}

}
