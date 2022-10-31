package main

import "fmt"

type OopPerson struct {
	// 使用结构体 加绑定函数实现类与方法
	name string
	age  int
}

func (op OopPerson) eat() {
	fmt.Println("oopPerson eating....")
}

func (op OopPerson) sleep() {
	fmt.Println("ooPerson sleeping....")
}

func (op OopPerson) work() {
	fmt.Println("ooPerson working......")
}

func OopTestDemo() {
	operson := OopPerson{"hewuxin", 24}

	fmt.Printf("per: %v\n", operson)
	operson.eat()
	operson.sleep()
	operson.work()
}
