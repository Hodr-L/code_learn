package chazhao

import "fmt"

func Erfen(arr *[6]int, leftindex int, rightindex int, findvalue int) {
	if leftindex > rightindex {
		fmt.Println("找不到！")
		return
	}
	middle := (leftindex + rightindex) / 2
	if (*arr)[middle] > findvalue {
		Erfen(arr, leftindex, middle-1, findvalue)
	} else if (*arr)[middle] < findvalue {
		Erfen(arr, middle+1, rightindex, findvalue)
	} else {
		fmt.Printf("找到了,下标为：%v 数据为：%v\n", middle, arr[middle])
	}
}
