package main

import "fmt"

func main() {
	i := 10
	var p = &i
	a := map[string]string{
		"Username": "admin",
		"Password": "admin",
	}
	//for _, value := range a {
	//fmt.Println(value)
	fmt.Println(&a)
	fmt.Println("--------")
	fmt.Println(&i)
	fmt.Println(*p)
	fmt.Println("--------")
	*p = 100
	fmt.Println(*p)
	fmt.Println(i)

	//}
}
