package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func ChannelTest() {
	//Unbuffered := make(chan int)      // 整型无缓冲通道
	buffered := make(chan string, 10) // 字符串型 10个缓冲的缓冲通道

	buffered <- "China"

	fmt.Println("buffered: ", buffered)
	//fmt.Println(Unbuffered)

	data := <- buffered
	fmt.Println("data: ", data)
}
*/

// 定义无缓冲通道
var values = make(chan int)

func send() {
	// 设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		value := rand.Intn(10) // 生成随机数

		fmt.Printf("send : %v \n", value)

		time.Sleep(time.Second)

		values <- value // 向通道写入value

	}

}

func fetch() {
	for i := 0; i < 5; i++ {
		value := <-values // 从values通道中获取值
		fmt.Printf("Recive: %v \n", value)
	}
}

func ChannelTest() {
	// 从通道接收值
	defer close(values) //当ChannelTest函数执行完成时关闭values通道
	go send()           // 开启协程执行send函数
	fmt.Println("wait...")

	go fetch()
	time.Sleep(time.Second * 3)
	fmt.Println("end...")
}
