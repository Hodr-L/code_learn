package array

import (
	"fmt"
	"math/rand"
	"time"
)

func Array04() {
	var b [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(100)
	}
	for i := 0; i < len(b); i++ {
		c := b[len(b)-1-i] //倒叙数组
		b[len(b)-1-i] = b[i]
		b[i] = c
	}
	fmt.Println(b)
}
