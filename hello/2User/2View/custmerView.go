package _View

import (
	"fmt"
	Modle "hello/2User/1Modle"
	Contral "hello/2User/3Contral"
)

type custmerView struct {
	//定义必要的字段
	key            string
	loop           bool
	custmerService *Contral.CustmerService
}

//得到用户的输入信息，构建新的用户，并完成添加
func (this *custmerView) add() {
	fmt.Println("---------添加客户-----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scan(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scan(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scan(&age)
	fmt.Println("手机号：")
	phone := 0
	fmt.Scan(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scan(&email)

	//构建一个新的Custmoer实例
	//注意：没有让用户输入id，因为id是唯一的，需要系统分配
	customer := Modle.NewCustmer2(
		name, gender, age, phone, email)
	if this.custmerService.Add(*customer) {
		fmt.Println("添加成功！")
	} else {
		fmt.Println("添加失败！")
	}
}

//得到用户的输入信息，删除对应id的用户
func (this *custmerView) delete() {
	fmt.Println("---------删除客户-----------")
	fmt.Println("请输入待删除的客户id（-1退出）：")
	id := -1
	fmt.Scan(&id)
	if id == -1 {
		return
		//放弃删除
	}
	for {
		fmt.Println("确认是否删除（Y/n）")
		cor := ""
		fmt.Scan(&cor)
		if cor == "Y" || cor == "y" {
			if this.custmerService.Delete(id) {
				fmt.Println("删除成功！")
				break
			} else {
				fmt.Println("删除失败,输入的id不存在！")
			}
		} else if cor == "N" || cor == "n" {
			this.delete()
		}
	}
}

func (this *custmerView) update() {
	fmt.Println("---------修改客户-----------")
	fmt.Println("请输入修改客户的id")
	id := 0
	fmt.Scan(&id)
	for {
		if this.custmerService.Delete(id) {
			fmt.Println("请输入要修改的信息！")
			break
		} else {
			fmt.Println("修改失败,输入的id不存在！")
			return
		}
	}
	fmt.Println("姓名：")
	name := ""
	fmt.Scan(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scan(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scan(&age)
	fmt.Println("手机号：")
	phone := 0
	fmt.Scan(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scan(&email)
	customer := Modle.NewCustmer2(
		name, gender, age, phone, email)
	if this.custmerService.Add(*customer) {
		fmt.Println("修改成功！")
	} else {
		fmt.Println("修改失败！")
	}
}

//显示所有客户信息
func (this *custmerView) list() {
	//首先要获取当前所有客户的信息
	cust := this.custmerService.List()
	//显示
	fmt.Println("---------客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(cust); i++ {
		fmt.Println(cust[i].GetInfo())
	}
	fmt.Println("---------客户列表完成-----------")
	fmt.Println()
}

func (cv *custmerView) mainView() {
	for {
		fmt.Println("---------客户信息管理软件----------")
		fmt.Println("1.添加客户")
		fmt.Println("2.修改客户")
		fmt.Println("3.删除客户")
		fmt.Println("4.客户列表")
		fmt.Println("5.退出")
		fmt.Println()
		fmt.Print("请选择(1-5):")
		fmt.Scan(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入有误，请重新输入！")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出！")
}

func CustmerView() {
	cust := custmerView{
		key:  "",
		loop: true,
	}
	//初始化
	cust.custmerService = Contral.NewcustService()
	//显示主菜单
	cust.mainView()
}
