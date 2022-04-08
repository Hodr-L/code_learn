package main

import (
	"fmt"
	"hello/demoDay23/goroutine"
	"sync"
	"time"
)

func main() {
	//go goroutine.Goroutine()
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("hello golang" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	//num := runtime.NumCPU()
	//fmt.Println(num)

	var lock sync.Mutex
	//开启多个协程完成任务
	for i := 1; i <= 20; i++ {
		go goroutine.Jiecheng1(i)
	}
	//主线程休眠，是为了让协程完成任务
	time.Sleep(5 * time.Second)
	//加锁为了让主线程知道协程运行完毕，防止资源竞争
	lock.Lock()
	for i, v := range goroutine.MyMap {
		fmt.Printf("第%d个元素是：%d\n", i, v)
	}
	lock.Unlock()
}
