package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type OcpPerson struct {
	name string
}

func (operson OcpPerson) care(pet Pet) {
	// 调用传入的结构体的绑定方法
	pet.eat()
	pet.sleep()
}

func ocpDemo() {
	dog := Dog{}
	cat := Cat{}
	person := OcpPerson{"hewuxin"}

	person.care(dog)
	person.care(cat)
}