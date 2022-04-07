package map01

import "fmt"

func Map03() {
	type user struct {
		Name   string
		Age    int
		Adress string
	}
	students := make(map[string]user)

	stu1 := user{"Tom", 19, "北京"}
	stu2 := user{"Jerry", 20, "大连"}
	students["No.1"] = stu1
	students["No.2"] = stu2

	for v, k := range students {
		fmt.Println(v, k.Name)
		fmt.Println(v, k.Age)
		fmt.Println(v, k.Adress)
	}
}
