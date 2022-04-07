package array

import "fmt"

func Array02() {
	var a float64
	var b [5]float64
	for i := 0; i < 5; i++ {
		fmt.Println("请输入成绩：")
		fmt.Scanln(&a)
		b[i] = a

	}
	{
		fmt.Println("输入结束！")
		fmt.Println()
		fmt.Println("输出结果为：")
	}
	for c := range b {
		fmt.Println(c)
	}
}
