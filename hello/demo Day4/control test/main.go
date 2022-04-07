package main

import "fmt"

func char(i byte) {
	switch i {
	case 'a':
		fmt.Printf("A")
	case 'b':
		fmt.Printf("B")
	case 'c':
		fmt.Printf("C")
	case 'd':
		fmt.Printf("D")
	case 'e':
		fmt.Printf("E")
	default:
		fmt.Printf("other!")

	}

}

func score(a float64) {
	switch {
	case a >= 60 && a <= 100:
		fmt.Printf("及格")
	case a < 60 && a >= 0:
		fmt.Printf("不及格")
	default:
		fmt.Printf("error!")
	}
}

func month(a int) {
	switch a {
	case 12, 1, 2:
		fmt.Printf("冬季！")
	case 3, 4, 5:
		fmt.Printf("春季！")
	case 6, 7, 8:
		fmt.Printf("夏季！")
	case 9, 10, 11:
		fmt.Printf("秋季！")
	default:
		fmt.Printf("error!")
	}
}

func time(a string) {
	switch a {
	case "星期一":
		fmt.Printf("干煸豆角")
	case "星期二":
		fmt.Printf("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	}
}

func main() {
	//var i byte
	//fmt.Println("请输入一个字符：")
	//fmt.Scanf("%c", &i)
	//char(i)
	//var score1 float64
	//fmt.Printf("请输入成绩：")
	//fmt.Scanf("%f", &score1)
	//score(score1)
	//var b int
	//fmt.Printf("请输入月份：")
	//_, _ = fmt.Scanf("%d", &b)
	//month(b)

	var eat string
	fmt.Println("请输入星期时间：")
	_, _ = fmt.Scanln(&eat)
	time(eat)
}
