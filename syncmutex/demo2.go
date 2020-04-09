package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	workers := 3
	wg.Add(workers)
	worker := func(i int) {
		defer wg.Done()
		// 干活干活干活
		print(i)
	}
	leader := func() {
		wg.Wait()
		// 检查工作成果
		print("检查工作成果")
	}
	go leader()
	for i := 0; i < workers; i++ {
		go worker(i)
	}
	time.Sleep(1)
}
