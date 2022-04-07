package testFunc

import "fmt"

//type Num struct {
//	Num int
//}
//func (a Num) Test() {
//	fmt.Printf("%v 是一个好人", a.Name)
//}
//func (a Num) Jisuan(b int, c int) {
//	a.Num = b + c
//	fmt.Println(a.Num)
//}

type Circle struct {
	Name string
	Age  int
}

func (a *Circle) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", a.Name, a.Age)
	return str
}
