package goroutine

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Goroutine() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}
}

var (
	MyMap = make(map[int]int, 10)
	//lock是一个全局互斥锁
	lock sync.Mutex
)

func Jiecheng1(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	MyMap[n] = res
	//解锁
	lock.Unlock()
}
