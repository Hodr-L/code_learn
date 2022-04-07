package face

import (
	"fmt"
)

//type Student struct {
//	Name   string
//	Gender string
//	Age    int
//	Id     int
//	Score  float64
//}
//
//func (student *Student) Say() string {
//	inforStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%d] id=[%d] score=[%f]",
//		student.Name, student.Gender, student.Age, student.Id, student.Score)
//	return inforStr
//}

//type Dog struct {
//	Name   string
//	Age    int
//	Weid float64
//}
//
//func (dog *Dog) Say() string {
//	inter := fmt.Sprintf("name=[%v] age=[%d] weight=[%.1f]", dog.Name, dog.Age, dog.Weid)
//	return inter
//}

//type Box struct {
//	Len  float64
//	Wide float64
//	High float64
//}
//func (box *Box) Say() float64 {
//	return box.Len * box.Wide * box.High
//}

type jingqu struct {
	Name string
	Age  int
}

func Newjingqu(n string, s int) *jingqu {
	return &jingqu{
		Name: n,
		Age:  s,
	}
}

func (jingqu *jingqu) Say() {

	if jingqu.Age >= 80 || jingqu.Age <= 5 {
		fmt.Println("为了您的安全，不建议游玩！")
		return
	}
	if jingqu.Age > 18 {
		fmt.Printf("%v的年龄为：%d,门票为20.\n", jingqu.Name, jingqu.Age)
	} else {
		fmt.Printf("%v的年龄为：%d,门票免费.\n", jingqu.Name, jingqu.Age)
	}
}
