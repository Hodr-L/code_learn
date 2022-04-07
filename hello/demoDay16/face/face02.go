package face

type student struct {
	name string
	age  int
}

func Stu(n string, a int) *student {
	return &student{
		name: n,
		age:  a,
	}
}
func (stu *student) Setname(n string) {
	stu.name = n
}
func (stu *student) Setage(a int) {
	stu.age = a
}

func (stu *student) Getname() string {
	return stu.name
}
func (stu *student) Getage() int {
	return stu.age
}
