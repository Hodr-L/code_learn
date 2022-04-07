package main

import "fmt"

func equal() bool {
	a := [10]string{"a", "b"}
	var x = &a
	var y = &a
	if x != y {
		return false
	}
	for i := range a {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func main() {
	c := equal()
	fmt.Println(c)
}
