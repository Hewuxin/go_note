package main

import (
	"fmt"
	"reflect"
)

/*
type 接口名称 interface{
	方法名称() 返回值
	接口中的方法只定义，不能编写具体的实现逻辑。
	type Base interface {
	f1()                   //定义方法，无返回值
	f2() int               //定义方法，返回值int类型
	f3() (int, bool)       // 定义方法，2个返回值int和bool
	f4(n1 int, n2 int) int // 输入参数 int int 返回值int
}

type empty interface{}
}
*/

func interfaceDemo() {
	//interDemo1()
	//interDemo2()
	//interDemo3()
	interDemo4()
}

// Base1 定义空接口
type Base1 interface{}

func interfaceDemo1() {
	//定义一个切片，内部可以存放任意类型.
	dataList := make([]Base1, 0)

	// 切片中添加 字符串类型
	dataList = append(dataList, "heyuyang")
	dataList = append(dataList, 11)
	dataList = append(dataList, 3.4)
	dataList = append(dataList, true)
	fmt.Println(dataList)
}

type interPerson struct {
	name string
	age  int
}

func someThing(arg interface{}) {

	fmt.Println(arg)
}

func interDemo1() {
	fmt.Println("This is interDemo1 接口声明")
	someThing("heyuyang")
	someThing(18)
	someThing("sex")
	someThing(true)
	someThing(4.15)
	someThing(interPerson{name: "heyuyang", age: 18})
}

// 由于接口只是代指这些数据类型(在内部其实是转换为了接口类型),
//想要再获取数据中的值时，需要再将接口转换为指定的数据类型。

func someThing1(arg interface{}) {

	// 接口转换为Person成功，ok=true, 否则ok=false
	fmt.Println("arg 的类型: ", reflect.TypeOf(arg))
	tp, ok := arg.(interPerson)
	fmt.Println("arg.(interPerson)之后的 arg的类型: ", reflect.TypeOf(arg))
	if ok {
		fmt.Println(tp.name, tp.age)
	} else {
		fmt.Println("转换失败")
	}
}

func interDemo2() {
	fmt.Println("This is interDemo2 接口转换")
	someThing1(interPerson{name: "何雨阳", age: 24})
}

type Role struct {
	title string
	count int
}

func something2(arg interface{}) {
	switch tp := arg.(type) {
	case interPerson:
		fmt.Println(tp.name)
	case Role:
		fmt.Println(tp.title)
	case string:
		fmt.Println(tp)
	default:
		fmt.Println(tp)
	}
}

func interDemo3() {
	something2(interPerson{"hewuxin", 24})
	something2(Role{"student", 18})
	something2("This is interDemo3")
	something2(true)
}

// IBase 定义接口 返回值int
type IBase interface {
	f1() int
}

// IPerson 定义结构体Person
type IPerson struct {
	name string
}

// 为结构体IPerson定义方法
func (p IPerson) f1() int {
	return 123
}

type User struct {
	name string
}

func (p User) f1() int {
	return 666
}

// InterSomeThing 基于接口的参数，可实现传入多种类型（多态）,也同时具有约束对象必须实现接口方法的功能
func InterSomeThing(base IBase) {
	result := base.f1() // 直接调用 接口.f1() -> 找到其对应的类型并执行其方法
	fmt.Println(result)
}

func interDemo4() {
	per := IPerson{"hewxuin"} // 创建结构体对象
	user := User{"student"}
	fmt.Println("per: ", per) // 调用对象的结构体方法
	InterSomeThing(per)
	fmt.Println("user: ", user)
	InterSomeThing(user)
}
