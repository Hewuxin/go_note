package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

func ErrorTest() {
	fmt.Println("errors包实现了一个最简单的error类型，只包含了一个字符串，他可以记录大多数情况下遇到的错误信息。")
	s, err := check("")
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
	} else {
		fmt.Printf("string: %v \n", s)
	}
}

// MyError is an error implementation that includes a time message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func ErrorTest01() {
	fmt.Println("自定义错误")
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
