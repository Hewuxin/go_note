package main

import (
	"fmt"
	"time"
)

// Timer 定时器，可以实现一些定时操作，定时器内部也是通过channel实现。

func NewTimerTest() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-timer.C // 阻塞的，指定时间到了
	fmt.Printf("t1 : %v \n", t1)
}

func NewTimerTest1() {
	fmt.Printf("time.Now(): %v \n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now())
}

func TimerAfterTest() {
	<-time.After(time.Second * 2)
	fmt.Println("......")
}

func TimerAfterTest1() {
	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("func....")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("main end...")
}

func TimerStop() {
	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("func......")
	}()

	s := timer.Stop()
	if s {
		fmt.Println("stop....")
	}
}

func TimerReset() {
	fmt.Println("before")
	timer := time.NewTimer(time.Second * 5) // 原来设置5s
	timer.Reset(time.Second * 1)            // Reset修改为1秒
	<-timer.C
	fmt.Println("After")
}

func TimerTicker() {
	fmt.Println("Timer 只执行一次，Ticker可以周期的执行.")
	ticker := time.NewTicker(time.Second)
	counter := 1

	for _ = range ticker.C {
		fmt.Println("ticker....")
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}

func TimerTicker1() {
	chanInt := make(chan int)

	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("接受: %v \n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("sum: %v\n", sum)
			break
		}
	}
}
