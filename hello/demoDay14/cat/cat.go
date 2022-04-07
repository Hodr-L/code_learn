package cat

import (
	"encoding/json"
	"fmt"
)

func Cat() {
	//var a string
	//var cat1 = []string{"小花", "100", "花色"}
	//var cat2 = []string{"小白", "3", "白色"}
	//fmt.Println("请输入：")
	//fmt.Scan(&a)
	//for _, i := range cat1 {
	//	if a == i {
	//		fmt.Println(cat1)
	//		return
	//	}
	//	for _, c := range cat2 {
	//		if a == c {
	//			fmt.Println(cat2)
	//			return
	//		} else {
	//			fmt.Println("小猫不存在!")
	//			return
	//		}
	//	}
	//}//
	//传统解决方法

	//结构体解决方法
	type Cat struct {
		Name  string `json:"name"` //结构体标签
		Age   int    `json:"age"`
		Color string `json:"color"`
		Hobby string `json:"hobby"`
	}

	var a string
	var cat1 Cat = Cat{
		Name: "小花",
	}
	cat1.Age = 5
	cat1.Color = "花色"
	cat1.Hobby = "<·)))><<"

	var cat2 Cat
	cat2.Name = "小白"
	cat2.Age = 3
	cat2.Color = "白色"
	cat2.Hobby = "<·)))><<"
	fmt.Println("请输入：")
	fmt.Scanln(&a)
	if a == cat1.Name {
		fmt.Println(cat1)
	} else if a == cat2.Name {
		fmt.Println(cat2)
	} else {
		fmt.Println("猫猫不存在！")
	}
	//var dog *Cat = new(Cat)
	//dog.Age = 1
	//fmt.Println(dog)
	//json.Marshak底层中用到了反射
	dog1, _ := json.Marshal(cat1)
	fmt.Println(string(dog1))
}
