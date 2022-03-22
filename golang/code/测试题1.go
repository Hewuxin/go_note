package main

import (
	"fmt"
	"math/rand"
	"time"
)

// bug1 生成的十条数据都一样
// 原因： 因为每次操作的间隔都在毫秒级下，
//       所以每次time.Now().Unix() 取出来的值都是同一个值， 就是使用了同一个seed
//       每次rand都会使用相同的seed来生成随机队列，这样一来在循环中使用相同seed得到的随机队列都是相同的，
//       而生成随机数时每次都会去取同一个位置的数，所以每次取到的随机数都是相同的
/*
seed 只用于决定一个确定的随机序列。不管seed多大多小，
只要随机序列一确定，本身就不会再重复。除非是样本空间太小。解决方案有两种：
	在全局初始化调用一次seed即可
	每次使用纳秒级别的种子（强烈不推荐这种）
*/

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("%6v  %6v   %5v    %5v\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("=================================================")
	for count := 0; count < 10; count++ {
		ticket_call()
		// time.Sleep(time.Second)
	}

}

func ticket_call() {
	// rand.Seed(time.Now().Unix())
	spaceline := Spaceline()
	days := Days()
	trip_type := Trip_type()
	price := Price()
	if trip_type == "Round-trip" {
		price = price * 2
	}
	fmt.Printf("%6v  %-6v   %-5v   %-5v\n", spaceline, days, trip_type, price)
}

func Price() int {
	// rand.Seed(time.Now().Unix())
	return rand.Intn(14) + 36
}

func Trip_type() string {
	// rand.Seed(time.Now().Unix())
	trip_num := rand.Intn(2) + 1
	res := "a"
	switch trip_num {
	case 1:
		res = "Round-trip"
	case 2:
		res = "One-Way"
	}
	return res
}

func Days() int {
	// rand.Seed(time.Now().Unix())
	speed := rand.Intn(14) + 16
	day := 621000000 / (speed * 60 * 60 * 24)
	return day
}

func Spaceline() string {
	// rand.Seed(time.Now().Unix())
	spaceline_id := rand.Intn(4) + 1
	spaceline := ""
	switch spaceline_id {
	case 1:
		spaceline = "SpaceAventures"
	case 2:
		spaceline = "SpaceX"
	case 3:
		spaceline = "Virgin Galactic"
	case 4:
		spaceline = "East Space"
	}
	return spaceline
}
