package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
func show(msg string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("msg: %v \n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func RoutineTest() {
	go show("JAva")
	show("golang")
}
*/

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Step2: ", url)

	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Step4: ", len(body))
}

func RoutineTest() {
	go responseSize("https://www.douban.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(10 * time.Second)
}
