package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 定义12个任务
	var quests = []string{"eat", "sleep", "drive", "shopping", "cooking", "running", "write", "read", "play", "walk", "dreaming", "watching"}
	// 建立容量为cpu数目的一般大小的缓冲信道
	var ch chan int = make(chan int, runtime.NumCPU()/2)

	for _, quest := range quests {
		go doQuestByChan(ch, quest)
	}

	// 等待全部goroutine结束
	for _ = range quests {
		<-ch
	}
}

// 做任务
func doQuestByChan(ch chan int, quest string) {
	fmt.Println(quest)

	// 代表任务已经完成
	ch <- 0
}
