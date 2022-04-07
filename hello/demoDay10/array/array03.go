package array

import "fmt"

func Array03() {
	var zimu [5]int = [...]int{1, 2, 5, 90, -5}
	max := zimu[0]
	c := 0
	for i := range zimu {
		if max <= zimu[i] {
			max = zimu[i]
			c = i
		}
	}
	fmt.Println(c, max)
}
