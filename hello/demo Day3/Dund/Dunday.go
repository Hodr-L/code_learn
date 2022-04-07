package main

import "fmt"

func test() bool {
	fmt.Println("test---------ok")
	return true
}
func test2() bool {
	fmt.Println("---------test")
	return true
}
func main() {
	a := 90
	b := 10
	if a > b || test() {
		fmt.Println("test----------")
	} else {
		test2()
	}

}
