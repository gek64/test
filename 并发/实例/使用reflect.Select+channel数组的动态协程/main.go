package main

import (
	"fmt"
	"reflect"
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
	// 按cpu数目设置信道数组
	var chs = make([]chan int, 2)

	// 信道建立并启动
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int)
		go func(ch chan int) {
			ch <- 1
		}(chs[i])
	}

	// 声明selectCase[]
	var selectCase = make([]reflect.SelectCase, len(chs))
	// 给每个selectCase 设置目的地和对应的信道
	for i := 0; i < len(selectCase); i++ {
		selectCase[i].Dir = reflect.SelectRecv
		selectCase[i].Chan = reflect.ValueOf(chs[i])
	}

	// 处理list中每一个
	for _, element := range list {
		// 使用reflect.Select 来在selectCase中选择可以用的信道
		// chosen为选择的信道序号,recv为目的地获取到的值,recvOk为是否成功选择
		chosen, recv, recvOk := reflect.Select(selectCase)

		if recvOk {
			fmt.Printf("chan %d is now available, the value case received is %v, it will call \"%s\"\n", chosen, recv, element)
			go call(element, chs[chosen])
		}
	}

	for _, ch := range chs {
		<-ch
	}
}

func call(str string, ch chan int) {
	fmt.Println(str)
	ch <- 1
}
