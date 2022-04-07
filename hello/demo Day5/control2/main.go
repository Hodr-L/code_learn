package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test1() {
	var num int
	fmt.Printf("乘法表的位数:")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d ", k, i, i*k)
		}
		fmt.Printf("\n")
	}
}
func test2() {
	var c int = 1

	for {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(100) + 1
		if a == 99 {
			fmt.Println(c)
			break
		}
		c++
	}
}
func test3() {
	var sum int
	for a := 1; a <= 100; a++ {
		sum += a
		if sum > 20 {
			fmt.Printf("%d", a)
			break
		}

	}
}
func test4() {
	var username string = "张无忌"
	var password string = "888"
	var a string
	var b string
	var i = 1
table2:
	for i <= 3 {
		for {
			fmt.Println("请输入用户名：")
			fmt.Scanln(&a)
			if a == username {
				fmt.Println("请输入密码：")
				fmt.Scanln(&b)
				if b == password {
					fmt.Println("登陆成功！")
					break table2
				} else {
					if 3-i > 0 {
						fmt.Printf("密码错误，还有%d次机会!", 3-i)
						i++
					} else {
						break table2
					}
				}
			} else {
				if 3-i > 0 {
					fmt.Printf("用户名错误，还有%d次机会!", 3-i)
					i++
				} else {
					break table2
				}
			}
		}
	}
}
func test5() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
func test6() {
	var a int
	fmt.Println("请输入：")
	var i int = 0
	var n int = 0
	for {
		fmt.Scanln(&a)
		if a == 0 {
			break
		} else if a > 0 {
			i++
			fmt.Println(i)
			continue
		}
		n++
		fmt.Println(n)
	}
}
func test7() {
	var a = 100.0
	for {
		var n int
		var k int
		if a == 0 {
			fmt.Println(n + k)
			break
		} else if a > 50 {
			a = a - (a * 0.05)
			fmt.Println(a)
			n++
		} else if a <= 50 && a > 0 {
			a = a - 10
			fmt.Println(a)
			k++
		}
	}
}

func main() {

}
