package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine! ", i)
}
func WaitGroupTest() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine 就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
*/

func showMsg(i int) {
	defer wg.Done() // 等价于 defer wg.Add(-1)
	time.Sleep(time.Second)
	fmt.Printf("i: %v \n", i)
}

func WaitGroupTest1() {
	for i := 0; i < 10; i++ {
		// 启动一个协程来执行
		wg.Add(1)
		go showMsg(i)
	}
	// 主协程
	fmt.Println("ending...")
	defer wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("end>>>>>>>>>>>>>.....")
}
