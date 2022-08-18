package main

import (
	"fmt"
	"time"
)

// 使用 flowControlChan 做流量控制
// 控制主进程发送协程的速率
//
// 使用 doneChan 做同步控制
// 控制主进程与协程同步完成

func main() {
	// 定义12个任务
	var quests = []string{"eat", "sleep", "drive", "shopping", "cooking", "running", "write", "read", "play", "walk", "dreaming", "watching"}

	// 建立容量为5的缓冲信道
	// 容量代表一次主进程发送运行的协程数目
	var flowControlChan = make(chan int, 5)

	// 建立容量为5的同步信道
	// 容量代表一次协程能发送已完成信号的数目
	var doneChan = make(chan bool, 5)

	for i, quest := range quests {
		// 流量控制信道(容量5)塞入一个标记表示已启动一个协程
		// 塞满之后要继续塞就得等待
		flowControlChan <- i

		// 启动协程
		go doQuestByChanWithFlowControl(flowControlChan, doneChan, quest)
	}

	// 同步控制,从同步信道中拿出所有的协程已完成信号
	for _ = range quests {
		<-doneChan
	}

}

// 做任务
func doQuestByChanWithFlowControl(flowControlChan chan int, doneChan chan bool, quest string) {
	// 任务内容
	fmt.Println(quest)
	time.Sleep(1 * time.Second)

	// 流量控制信道中取出标记
	<-flowControlChan

	// 同步信道中发送已完成信号
	doneChan <- true
}
