package Modle

import "fmt"

//声明一个custmer结构体，表示一个客户信息

type Custmer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  int
	Email  string
}

func NewCustmer(id int, name string, gender string, age int, phone int, email string) *Custmer {
	return &Custmer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func NewCustmer2(name string, gender string, age int, phone int, email string) *Custmer {
	return &Custmer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回客户的信息,格式化字符串

func (this Custmer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
