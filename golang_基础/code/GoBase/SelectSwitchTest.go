package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int) //
var chanStr1 = make(chan string)

func SelectTest() {
	go func() {
		chanInt <- 100      // 将100写入通道
		chanStr1 <- "hello" // 将hello写入通道
		close(chanInt)
		close(chanStr1)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr1:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Println("default....")
		}
		time.Sleep(time.Second)
	}

}
