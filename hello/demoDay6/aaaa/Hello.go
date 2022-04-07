package aaaa

func Hello(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Hello(n-1) + Hello(n-2)
	}
}
func Res(n int) (sum int) {
	sum = 2*(n-1) + 1
	if n == 1 {
		return 3
	} else {

		return
	}
}
func Monkey(n int) int {
	if n == 10 {
		return 1
	} else {
		return (Monkey(n+1) + 1) * 2
	}
}
