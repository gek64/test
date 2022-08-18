package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义12个任务
	var quests = []string{"eat", "sleep", "drive", "shopping", "cooking", "running", "write", "read", "play", "walk", "dreaming", "watching"}
	// 新建同步序列
	var wg = sync.WaitGroup{}

	for _, quest := range quests {
		// 同步序列填充一个任务
		wg.Add(1)
		go doQuestByWG(&wg, quest)
	}

	// 等待所有同步序列任务完成
	wg.Wait()
}

// 做任务
func doQuestByWG(wg *sync.WaitGroup, quest string) {
	fmt.Println(quest)

	// 同步序列1个任务完成
	wg.Done()
}
