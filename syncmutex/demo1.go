package main

import "time"

func main() {
	workers := 3
	ch := make(chan struct{})
	worker := func(i int) {
		// 干活干活干活
		print(i)
		ch <- struct{}{} // 通知管理者
	}
	leader := func() {
		cnt := 0
		for range ch {
			cnt++
			if cnt == workers {
				break
			}
		}
		close(ch)
		// 检查工作成果
		print("over")
	}
	go leader()
	for i := 0; i < workers; i++ {
		go worker(i)
	}
	time.Sleep(10)
}
