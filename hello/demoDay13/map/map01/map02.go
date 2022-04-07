package map01

import (
	"fmt"
	"sort"
)

func Map02() {
	map02 := make(map[int]int)
	map02[10] = 100
	map02[1] = 13
	map02[8] = 16
	map02[4] = 19
	//fmt.Println(map02)

	var keys []int
	for k, _ := range map02 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//fmt.Println(keys)
	map02[4] = 200
	for _, v := range keys {
		fmt.Print("\n", v, map02[v])
	}
}
