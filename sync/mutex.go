package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	type mutex struct{
		state int32 //锁状态
		sema int32  //信号量 唤醒用
	}
	mutex锁
	如果新到的goroutine很多且正处于运行态，所以waiter等待者队列很难抢到🔒
	如果一个waiter等待了超过1ms，mutex直接变成饥饿状态，锁给队头
	如果waiter发现(1)它是队尾(2)等待时间少于1ms，mutex变回普通状态
	const(
		mutexLocked = 1<<iota //加锁状态
		mutexWoken			  //表示锁唤醒
		mutexStarving		  //表示锁是饥饿状态
		mutextWaiterShift = iota //表示waiter的个数
		starvationThresholdNs = 1e6 //1ms
	)
*/

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	wg.Add(2)
	go read()
	go read()
	wg.Wait()
	fmt.Println("finfish")
}
func read() {
	rwMutex.RLock()
	fmt.Println("read")
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	defer wg.Done()
}
