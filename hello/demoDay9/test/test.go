package test

import "fmt"

func Test() {
	var y string
	var m int
	var d int
	for {
		fmt.Println("请输入年份：")
		fmt.Scan(&y)
		fmt.Println("请输入月份：")
		fmt.Scan(&m)
		fmt.Println("请输入日期：")
		fmt.Scan(&d)

		if m <= 12 && m > 0 {
			fmt.Println("日期为：", y, "年", m, "月", d, "日")
		} else {
			fmt.Println("输入错误！")
		}

	}
}
