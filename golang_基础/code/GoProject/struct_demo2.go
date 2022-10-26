package main

import "fmt"

func StructDemo2() {
	fmt.Println("This is structDemo2")
	structDemo01()
	structDemo02()
	structDemo03()
}

type Person struct {
	name string
	age  int
}

var P Person = Person{"heyuayng", 18}

func doSomething() Person {
	return P
}

func structDemo01() {
	fmt.Println("使用结构体作为参数和返回值时，结构体作为返回值返回时，会拷贝一份因此修改原值不会修改返回值")
	data := doSomething()
	fmt.Println(data)
	P.name = "hewuxin"
	fmt.Println("data: ", data)
	fmt.Println("P: ", P)
}

func doSomething01() *Person {
	return &P
}

func structDemo02() {
	fmt.Println("使用结构体指针作为返回值")
	data := doSomething01()
	fmt.Println(data)
	P.name = "hewuxin"
	fmt.Println("data: ", data)
	fmt.Println("P: ", P)
}

type MyInt int

func (i *MyInt) DoSomething(a1 int, a2 int) int {
	// 为MyInt类型自定义一个指针方法
	// 可以是指针/可以是类型： *MyInt MyInt
	return a1 + a2 + int(*i)
}

func Do(a1 int, a2 int) int {
	return a1 + a2
}

func (_ *MyInt) DoSomething01(a1 int, a2 int) int {
	// 为MyInt类型自定义一个指针方法
	// 可以是指针/可以是类型： *MyInt MyInt
	return a1 + a2
}

func structDemo03() {
	fmt.Println("项目开发中可以为type声明的类型编写一些方法，从而实现 对象.方法的操作。")
	var v1 MyInt = 1
	result := v1.DoSomething(1, 2)
	fmt.Println("result: ", result)

	var v2 MyInt = 2
	res := v2.DoSomething01(2, 3)
	fmt.Println("res: ", res)
}

type People struct {
	name string
	age  int
	blog string
}

// 为People结构体类型自定义一个指针方法
// 注意：此处如果不是指针类型的话， 再执行方法时结构体对象就会被拷贝一份。
func (P *People) DoSomething02(a1 int, a2 int) int {
	return a1 + a2 + P.age
}

func structDemo04() {
	fmt.Println("注意：在方法名之前，func关键字之后的括号中指定receiver。如果方法不需要使用recv的值，可以使用_替换它。" +
		"recv就像是面向对象语言中的this或self，但是Go中并没有这两个关键字。可以使用this或self作为receiver的名字")
	p1 := People{name: "heyuyang", age: 18, blog: "https://www.pythonav.com"}
	result := p1.DoSomething02(1, 2)
	fmt.Println(result)
}

// 方法继承
// 如果结构体之前存在匿名嵌套关系，则子结构体可以继承父结构中的方法。

type Base struct {
	name string
}

type Son struct {
	Base // 匿名的方式，如果改成base Base 则无法继承Base的方法。
	age  int
}

func (b *Base) m1() int {
	return 666
}

func (s *Son) m2() int {
	return 9999
}

func structDemo05() {
	fmt.Println("This is StructDemo05 ")
	son := Son{age: 18, Base: Base{name: "heyuyang"}}
	result1 := son.m1()
	result2 := son.m2()
	fmt.Println(result1, result2)
}

type File1 struct {
	fd   int
	name string
}

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	return &File{fd, name}
}

func structDemo06() {
	fmt.Println("This is StructDemo06")
	message := " Go语言不支持面向对象编程的构造方法，但是可以很容易的在Go中实现“构造方法”。为了方便通常会为类型定义一个工厂，" +
		"按惯例，工厂的名字以New或new开头"
	fmt.Println("This is StructDemo06\n", message)
	f := NewFile(10, "./test.txt")
	fmt.Println(f.fd, f.name)
	f1 := File{10, "XXXXXX"}
	fmt.Println(f1.fd, f1.name)

}
