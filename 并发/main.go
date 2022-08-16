package main

import (
	"fmt"
	"reflect"
)

func main() {
	var completed chan int = make(chan int)

	go loop(10, completed)
	loop(5)

	<-completed

}

// completed 数组第一个参数 是 chan int 类型
func loop(n int, completed ...interface{}) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
	// 如果第二个参数传入 int 信道,则执行消息传入信道
	if len(completed) != 0 {
		if reflect.ValueOf(completed[0]).Kind() == reflect.Chan {
			completed[0].(chan int) <- 1
		}
	}
}
