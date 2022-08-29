package main

import (
	"fmt"
	"sync"
	"time"
)

// 可以做流量控制,但协程的无法用管道返回数据

// 使用 flowControlChan 做流量控制
// 控制主进程发送协程的速率
//
// 使用 sync.WaitGroup{} 做同步控制
// 控制主进程与协程同步完成

func main() {
	// 定义12个任务
	var quests = []string{"eat", "sleep", "drive", "shopping", "cooking", "running", "write", "read", "play", "walk", "dreaming", "watching"}

	// 建立容量为5的缓冲信道
	// 容量代表一次主进程发送运行的协程数目
	var flowControlChan = make(chan int, 5)

	// 建立同步组
	var wg = sync.WaitGroup{}

	for i, quest := range quests {
		// 流量控制信道(容量5)塞入一个标记表示已启动一个协程
		// 塞满之后要继续塞就得等待
		flowControlChan <- i

		// 同步控制计数器+1
		wg.Add(1)

		// 启动协程
		go doQuestByChanAndSyncWithFlowControl(flowControlChan, &wg, quest)
	}

	// 等待同步计数器归零
	wg.Wait()
}

// 做任务
func doQuestByChanAndSyncWithFlowControl(flowControlChan chan int, wg *sync.WaitGroup, quest string) {
	// 任务内容
	fmt.Println(quest)
	time.Sleep(1 * time.Second)

	// 流量控制信道中取出标记
	<-flowControlChan

	// 同步控制计数器-1
	wg.Done()
}
