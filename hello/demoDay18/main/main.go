package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}
type Camer struct {
}
type Computer struct {
}

func (phone Phone) Start() {
	fmt.Println("开始！")
}
func (phone Phone) Stop() {
	fmt.Println("停止！")
}
func (camer Camer) Start() {
	fmt.Println("开始~！")
}
func (camer Camer) Stop() {
	fmt.Println("停止~！")
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}
func (phone Phone) Call() {
	fmt.Println("0.0")
}

//type A interface {
//	Test01()
//}
//type B interface {
//	Test02()
//}
//type C interface {
//	A
//	B
//}
//type D struct {
//}
//
//func (d D) Test01() {
//}
//

func main() {
	camer := Camer{}
	phone := Phone{}
	com := Computer{}
	com.Working(phone)
	com.Working(camer)

}

//type Hero struct {
//	Name string
//	Age  int
//}
//type HeroSlice []Hero
//
//func (hs HeroSlice) Len() int {
//	return len(hs)
//}
//func (hs HeroSlice) Less(i, j int) bool {
//	//Less方法决定使用什么方法进行排序
//	return hs[i].Age > hs[j].Age
//}
//func (hs HeroSlice) Swap(i, j int) {
//	temp := hs[i]
//	hs[i] = hs[j]
//	hs[j] = temp
//}
//
//type Stu struct {
//	Name string
//	Id   int
//}
//
//type StuSlice []Stu
//
//func (stu StuSlice) Len() int {
//	return len(stu)
//}
//func (stu StuSlice) Less(i, j int) bool {
//	return stu[i].Id < stu[j].Id
//}
//func (stu StuSlice) Swap(i, j int) {
//	stu[i], stu[j] = stu[j], stu[i]
//}
//
//func main() {
//	//var intSlice = []int{0, -1, 10, 7, 90}
//	//sort.Ints(intSlice)
//	//fmt.Println(intSlice)
//
//	//var heros HeroSlice
//	//for i := 0; i < 10; i++ {
//	//	hero := Hero{
//	//		Name: fmt.Sprintf("张三~%d", rand.Intn(100)),
//	//		Age:  rand.Intn(100),
//	//	}
//	//	heros = append(heros, hero)
//	//}
//	//sort.Sort(heros)
//	//for _, v := range heros {
//	//	fmt.Println(v)
//	//}
//
//	var stu StuSlice
//	for i := 0; i < 10; i++ {
//		stu1 := Stu{
//			Name: fmt.Sprintf("李四~%d", rand.Intn(50)),
//			Id:   rand.Intn(100),
//		}
//		stu = append(stu, stu1)
//	}
//	sort.Sort(stu)
//	for _, b := range stu {
//		fmt.Println(b)
//	}
//
//}
