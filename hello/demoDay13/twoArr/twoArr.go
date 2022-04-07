package twoArr

import "fmt"

func TwoArr() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[j]); j++ {
	//		fmt.Print(arr[i][j])
	//	}
	//	fmt.Println()
	//}
	for _, v := range arr {
		for _, b := range v {
			fmt.Printf("%v", b)
		}
		fmt.Println()
	}
}
