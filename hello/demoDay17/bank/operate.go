package bank

import "fmt"

func Operate() {
	operate := Account{
		AccountNum: "admin",
		Password:   "1",
		Blance:     1200,
	}
	var a int
	var pwd string
	var money float64

	fmt.Println("请输入账户：")
	_, _ = fmt.Scan(&operate.AccountNum)
c:
	fmt.Println("请选择操作:")
	fmt.Println("1.存款")
	fmt.Println("2.取款")
	fmt.Println("3.查询余额")
	fmt.Println("4.退出系统")
	fmt.Print("输入：")
	_, _ = fmt.Scan(&a)
	for {
		switch {
		case a == 1:
			fmt.Println("请输入密码：")
			_, _ = fmt.Scan(&pwd)
			fmt.Println("请输入存钱数：")
			_, _ = fmt.Scan(&money)
			operate.Deposite(pwd, money)
			goto c
		case a == 2:
			fmt.Println("请输入密码：")
			_, _ = fmt.Scan(&pwd)
			fmt.Println("请输入取钱数：")
			_, _ = fmt.Scan(&money)
			operate.Withdrawal(pwd, money)
			goto c
		case a == 3:
			fmt.Println("请输入密码：")
			_, _ = fmt.Scan(&pwd)
			operate.Inquire(pwd)
			goto c
		case a == 4:
			break
		}
		break
	}
}
