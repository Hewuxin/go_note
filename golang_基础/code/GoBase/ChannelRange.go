package main

import "fmt"

func ChannelTest01() {
	c := make(chan int) // 注册无缓冲通道

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("将", i, "写入通道")
			c <- i // 向通道写入数据
		}
		close(c) // for循环写入完毕时关闭通道
	}()

	for {
		// 从通道中读取数据  死循环读取 如果读到值就打印，不然就跳出循环
		if data, ok := <-c; ok {
			fmt.Printf("data : %v \n", data)
		} else {
			break
		}
	}
}

func ChannelTest02() {
	c := make(chan int)

	go func() {
		for i := 0; i < 2; i++ { //
			fmt.Println("将", i, "写入通道")
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 3; i++ {
		r := <-c //  如果int通道中没有值 会读到默认值0
		fmt.Printf("r: %v \n", r)
	}
}

func ChannelTest03() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Printf("v: %v\n", v)
	}
}
