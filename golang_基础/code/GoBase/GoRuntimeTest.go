package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}
func GoschedTest() {
	go show("Golang")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 有权利执行任务，让给其他子协程执行
	}
	fmt.Println("ending...")
}
*/
/*
func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)

		if i >= 5 {
			runtime.Goexit()
		}
	}

}

func GoexitTest() {

	go show()

	time.Sleep(time.Second)
}
*/

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A: ", i)
	}
}
func a1() {
	for i := 0; i < 10; i++ {
		fmt.Println("a: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B: ", i)
	}
}

func b1() {
	for i := 0; i < 10; i++ {
		fmt.Println("b: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func GoMaxProcsTest() {
	fmt.Printf("runtime.NumCPU(): %v \n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go a()
	go a1()
	go b()
	go b1()
	time.Sleep(time.Second)
}
