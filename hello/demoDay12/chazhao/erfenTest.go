package chazhao

import (
	"fmt"
)

func ErfenTest(arr *[]int, leftindex int, rightindex int, find int) {
	middle := (leftindex + rightindex) / 2
	if leftindex > rightindex {
		fmt.Println("False")
		return
	}
	if find > (*arr)[middle] {
		ErfenTest(arr, middle+1, rightindex, find)
	} else if find < (*arr)[middle] {
		ErfenTest(arr, leftindex, middle-1, find)
	} else {
		fmt.Printf("True")
	}
}
