package bank

import "fmt"

type Account struct {
	AccountNum string
	Password   string
	Blance     float64
}

//1.存款
func (account *Account) Deposite(pwd string, money float64) {
	if pwd != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	if money <= 0 {
		fmt.Println("输入金额不正确！")
		return
	}
	account.Blance += money
	fmt.Println("存款成功！")
	fmt.Println("当前余额为：", account.Blance)
	return
}

//2.取款
func (account *Account) Withdrawal(pwd string, money float64) {
	if pwd != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	if money <= 0 || money > account.Blance {
		fmt.Println("输入金额不正确！")
		return
	}
	account.Blance -= money
	fmt.Println("取款成功！")
	fmt.Println("当前余额为：", account.Blance)
	return
}

//3.查询余额
func (account *Account) Inquire(pwd string) {
	if pwd != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	fmt.Printf("余额：%.2f\n", account.Blance)
	return
}
