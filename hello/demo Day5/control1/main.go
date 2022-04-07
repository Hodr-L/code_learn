package main

import (
	"fmt"
)

func con() {
	var i = 1
	for {
		if i <= 10 {
			fmt.Println("你好，World！", i)
			fmt.Println("**********")

		} else {
			break
		}
		i++
	}
}
func str() {
	var ab = "hello wprld!北京"
	//cd := []rune(ab)
	for i, a := range ab {
		fmt.Printf("%d\n 字符为：%c ", i, a)
	}
}
func test1() {
	var cont = 0
	var sum = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			cont++
			sum += i
		}
	}
	fmt.Println(cont, sum)
}
func test2(a int) {
	var b = 0
	var c = a
	for i := 0; i <= a+1; i++ {
		if b+c == a {
			if c >= 0 {
				fmt.Println(b, "+", c, "=", a)
			}
			b++
			c--
		}
	}
}
func test3() {
	var i = 0
	for {
		if i < 10 {
			fmt.Println("hello world!", i)
			i++
		} else {
			break
		}
	}
	fmt.Println(i)
}
func test4() {
	var b float64
	var sum float64
	var summ float64
	for a := 1; a <= 3; a++ {
		sum = 0
		for i := 1; i <= 5; i++ {
			fmt.Printf("请输入第%d个班第%d个学生的成绩：", a, i)
			fmt.Scanln(&b)
			sum += b
			summ += b
		}
		fmt.Printf("第%d个班的平均成绩为：%f\n", a, sum/5)
	}
	fmt.Printf("总平均分为：%f", summ/15.0)
}
func test5() {
	var scoce int = 0.0
	var d int = 0
	for a := 1; a <= 3; a++ {
		for i := 1; i <= 5; i++ {
			fmt.Printf("请输入第%d个班第%d同学的成绩：", a, i)
			fmt.Scanln(&scoce)
			if scoce >= 60 {
				d++
			}
		}
	}
	fmt.Printf("及格的人数为:%d人\n", d)
}
func test6() {
	var a int
	fmt.Println("请输入层数:")
	fmt.Scanf("%d", &a)
	for i := 1; i <= a; i++ {
		//for k := 1; k <= (a - i); k++ {
		//	fmt.Printf(" ")
		//}
		for c := 1; c <= (2*i)-1; c++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
func test7() {
	var a int
	fmt.Printf("请输入层数：")
	fmt.Scanf("%d", &a)

	for i := 1; i <= a; i++ {
		for k := 1; k <= (a - i); k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i)-1; j++ {
			if j == 1 || j == (2*i-1) || i == a {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func main() {
	test7()
}
