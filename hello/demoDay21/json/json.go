package json

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64 //薪水
	Skill    string  //技能
}

func TestSturst() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2020-2-31",
		Sal:      1500.2,
		Skill:    "牛魔拳",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败！")
	} else {
		fmt.Println(string(data))
	}
}

func TestMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 100
	a["adress"] = "火云洞"
	a["sal"] = 1500.23
	a["skill"] = "吃唐僧肉"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化map错误！")
	} else {
		fmt.Println(string(data))
	}
}

func TestSlice() {
	var a []map[string]interface{}
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "李云龙"
	m1["age"] = 53
	m1["adress"] = "山东"
	//m1["sal"] = 10000.5
	//m1["skill"] = "意大利炮"
	a = append(a, m1)

	m2 = make(map[string]interface{})
	m2["name"] = "楚云飞"
	m2["age"] = 54
	m2["adress"] = "台湾"
	//m2["sal"] = 1005.65
	//m2["skill"] = "日你仙人"
	a = append(a, m2)

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败！")
	} else {
		fmt.Println(string(data))
	}

}

func TestFloat64() {
	var num1 float64 = 2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败！")
	} else {
		fmt.Println(string(data))
	}
}
