package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eating....")
}

func (a Animal) sleep() {
	fmt.Println("sleeping...")
}

type SDog struct {
	Animal
	action string
}

func (sdog SDog) act() {
	fmt.Println("dog act.....")
}

type SCat struct {
	Animal
	fan string
}

func (scat SCat) jump() {
	fmt.Println("cat jumping....")
}

func superTest() {
	fmt.Println("golang通过结构体嵌套实现继承.")
	sd := SDog{Animal{"xiaohuang", 2}, "wangwang"}
	sd.eat() // Animal -> SDog
	sd.sleep()
	sd.act() // SDog自己绑定的方法

	sc := SCat{Animal{"tom", 3}, "jump"}
	sc.eat() // Animal -> Scat
	sc.sleep()
	sc.jump() // SCat自己绑定的方法
}

