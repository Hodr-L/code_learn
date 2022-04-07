package maopao

import "fmt"

func Maopao() {
	var max int
	arr := []int{15, 20, 80, 13, 60, 1, 2, 3, 4, 5, 7, 62, 745, 574}
	for a := 0; a < len(arr)-1; a++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				max = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = max
			}
		}
	}
	fmt.Println(arr)
}
