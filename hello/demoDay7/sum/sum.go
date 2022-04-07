package sum

import "fmt"

func Sum(args ...int) (num int) {
	for i := 0; i < len(args); i++ {
		num += args[i]
	}
	return num
}

func Sum1(n1, n2 float32) float32 {
	fmt.Printf("%T", n1)
	return n1 + n2
}
func Sum2(n1, n2 *int) {

	t := *n1
	*n1 = *n2
	*n2 = t

}
