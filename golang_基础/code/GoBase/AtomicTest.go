package main

import (
	"fmt"
	"sync/atomic"
)

/*
atomic提供的原子操作能够确保任一时刻只有一个Goroutine对变量进行操作，
善用atomic能够避免程序中出现大量的锁操作。
atomic常见操作：
- 增减
-载入read
-比较交换 cas
-交换 swap
- 存储 write
*/
//var i = 100
//
//var lock sync.Mutex
//
//func add() {
//	lock.Lock()
//	i++
//	lock.Unlock()
//}
//
//func sub() {
//	lock.Lock()
//	i--
//	lock.Unlock()
//}
//
//func AtomicTest01() {
//	for i := 0; i < 100; i++ {
//		go add()
//		go sub()
//	}
//
//	time.Sleep(time.Second * 3) //  sleep3秒防止主协程运行完毕直接杀死其他协程
//	fmt.Printf("i: %v\n", i)
//
//}

//var i int32 = 100
//
//// cas compare and swap
//func add1() {
//	atomic.AddInt32(&i, 1) // 跨作用域修改
//}
//func sub1() {
//	atomic.AddInt32(&i, -1) // 跨作用域修改
//}
//
//func AtomicTest02() {
//	for i := 0; i < 100; i++ {
//		go add1()
//		go sub1()
//	}
//	time.Sleep(time.Second * 3) //  sleep3秒防止主协程运行完毕直接杀死其他协程
//	fmt.Printf("i: %v\n", i)
//}

func LoadStoreTest() {
	fmt.Println("LoadStoreTest..........")
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Printf("i : %v \n", i)

	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

func CasTest() {
	// cas
	fmt.Println("CompareAndSwapInt32......")
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b : %v\n", b)
	fmt.Printf("i : %v\n", i)
}
