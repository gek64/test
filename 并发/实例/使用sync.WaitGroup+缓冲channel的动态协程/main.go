package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var list []string = []string{
		"this is 0",
		"this is 1",
		"this is 2",
		"this is 3",
		"this is 4",
		"this is 5",
	}
	// 按cpu数目设置缓冲channel
	var ch chan int = make(chan int, runtime.NumCPU())

	// 等待全部goroutine的计数器
	var wg = sync.WaitGroup{}

	// 处理list中每一个
	for _, element := range list {
		ch <- 1
		wg.Add(1)
		fmt.Printf("it will call \"%s\"\n", element)
		go call(element, ch, &wg)
	}

	// 等待全部goroutine结束后继续
	wg.Wait()
}

func call(str string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(str)
	<-ch
	wg.Done()
}
