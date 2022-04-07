package menu

import (
	"fmt"
)

type User struct {
	loop    bool
	a       string
	balance float64
	money   float64
	note    string
	details string
}

func A() *User {
	return &User{
		balance: 0,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
		loop:    true,
	}
}

func (user *User) call1() {
	fmt.Println("输入错误，请重新输入！")
	fmt.Println("-----------当前收支明细------------")
	if user.money == 0 {
		fmt.Println("当前没有收支明细，增加一笔吧！")
	} else {
		fmt.Println(user.details)
	}
}

func (user *User) call2() {
	fmt.Println("本次收入金额：")
	fmt.Scan(&user.money)
	user.balance += user.money
	fmt.Println("本次收入说明:")
	fmt.Scan(&user.note)
	user.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", user.balance, user.money, user.note)
}

func (user *User) call3() {
	fmt.Println("本次支出金额：")
	fmt.Scan(&user.money)
	if user.money > user.balance {
		fmt.Println("超出余额！")
	}
	user.balance -= user.money
	fmt.Println("本次支出说明:")
	fmt.Scan(&user.note)
	user.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", user.balance, user.money, user.note)
}

func (user *User) exit() {
	for {
		fmt.Println("你确定要退出吗？y/n")
		fmt.Scan(&user.a)
		if user.a == "y" {
			user.loop = false
		} else if user.a == "n" {
			user.Menu()
		} else {
			fmt.Println("输入错误！")
		}
		return
	}

}

func (user User) Menu() {
	var b User
	for {
		fmt.Println("--------记账薄--------")
		fmt.Println("1.查看明细")
		fmt.Println("2.收入")
		fmt.Println("3.支出")
		fmt.Println("4.退出")
		fmt.Print("请输入：")
		fmt.Scan(&b.a)
		fmt.Println()
		fmt.Println("--------记账薄--------")
		switch b.a {
		case "1":
			b.call1()
		case "2":
			b.call2()
		case "3":
			b.call3()
		case "4":
			b.exit()
		default:
			fmt.Println("输入错误，请重新输入！")
		}
		if !b.loop {
			break
		}
	}
}
