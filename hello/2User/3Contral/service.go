package Contral

import (
	"hello/2User/1Modle"
)

//该CustmerService完成对Custmer的操作，完成增删改查的操作
type CustmerService struct {
	customers []Modle.Custmer
	//声明一个字段，表示当前切片含有多少个客户
	custnum int
	//该字段可以作为新客户的（id+1）编号
}

func NewcustService() *CustmerService {
	cust := &CustmerService{}
	cust.custnum = 1
	customer := Modle.NewCustmer(
		1, "张三", "男", 18,
		12345678987, "dsadasd@163.com")
	cust.customers = append(cust.customers, *customer)
	return cust
}

func (this *CustmerService) List() []Modle.Custmer {
	return this.customers
}

//添加客户到custmer切片中
func (this *CustmerService) Add(custmer Modle.Custmer) bool {
	this.custnum++
	custmer.Id = this.custnum
	this.customers = append(this.customers, custmer)
	return true
}

func (this *CustmerService) Delete(id int) bool {
	index := this.FindId(id)
	if index == -1 {
		return false
	}
	//如何从切片中删除一个客户
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
func (this *CustmerService) Update(custmer Modle.Custmer) {
	this.FindId(this.custnum)
	this.Delete(this.custnum)
	this.customers = append(this.customers, custmer)
}
func (this *CustmerService) FindId(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
