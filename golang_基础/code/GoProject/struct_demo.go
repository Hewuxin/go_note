package main

import (
	"fmt"
	"reflect"
)

func structDemo0() {
	fmt.Println("This is structDemo -- 什么是结构体")
	type Person struct {
		name  string
		age   int
		email string
	}

	var p1 = Person{"heyuyang", 19, "daxinzang@163.com"}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age, p1.email)

	p1.age = 20
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age, p1.email)
}

func structDemo1() {
	fmt.Println("This is structDemo1 -- 结构体初始化")
	type Person struct {
		name  string
		age   int
		hobby []string
	}
	fmt.Println("先后顺序定义初始化")
	var p1 = Person{"hewuxin", 19, []string{"ba", "soc"}}
	fmt.Println(p1.name, p1.age)
	for _, item := range p1.hobby {
		fmt.Println(item)
	}

	fmt.Println("关键字初始化")
	var p2 = Person{name: "hewuxin", age: 19, hobby: []string{"ba", "soc"}}
	fmt.Println(p2.name, p2.age)
	for _, item := range p2.hobby {
		fmt.Println(item)
	}

	fmt.Println("先声明再赋值")
	var p3 Person
	p3.name = "dasdasd"
	p3.age = 120
	p3.hobby = []string{"a", "b"}
	fmt.Println(p3.name, p3.age)
	for _, item := range p3.hobby {
		fmt.Println(item)
	}

}

func structDemo2() {
	fmt.Println("This is structDemo2 -- 嵌套结构体声明、初始化")
	type Address struct {
		city, state string
	}

	type Person struct {
		name string
		age  int
		ad   Address
	}

	fmt.Println("先后顺序定义初始化")
	var p1 = Person{"hewuxin", 19, Address{"beijing", "china"}}
	fmt.Println(p1.name, p1.age, p1.ad.city, p1.ad.state)

	fmt.Println("关键字初始化")
	var p2 = Person{name: "hewuxin", age: 19, ad: Address{"beijing", "china"}}
	fmt.Println(p2.name, p2.age, p2.ad.city, p2.ad.state)

	fmt.Println("先声明再赋值")
	var p3 Person
	var p4 Address
	p3.name = "dasdasd"
	p3.age = 120
	p3.ad = p4
	p4.city = "beijing"
	p4.state = "china"
	fmt.Println("p3: ", p3.name, p3.age, p3.ad.city, p3.ad.state)
	fmt.Println("p4 : ", p4.city, p4.state)
}

func structDemo3() {
	fmt.Println("This is structDemo3 -- 嵌套匿名字段结构体")
	type Address struct {
		city, state string
	}

	type Person struct {
		name    string
		age     int
		Address // 匿名字段，那么默认Person就包含了Address的所有字段
	}

	//方式1：先后顺序
	p1 := Person{"何雨阳", 19, Address{"北京", "china"}}
	fmt.Println(p1.name, p1.age, p1.city, p1.state)

	//方式2：关键字
	var p2 = Person{name: "hewuxin", age: 19, Address: Address{"北京", "china"}}
	fmt.Println(p2.name, p2.age, p2.Address.city, p2.Address.state, p2.Address.city, p1.Address.state)

	//方式3：先声明再赋值
	var p3 Person
	p3.name = "hewuxin"
	p3.age = 18
	p3.Address = Address{
		city:  "beijing",
		state: "china",
	}

	fmt.Println(p3.name, p3.age, p3.Address.city, p3.Address.state)
	// 或
	var p4 Person
	p4.name = "hewuxin"
	p4.age = 18
	p4.city = "beijing"
	p4.state = "china"
	fmt.Println("先声明再赋值")
	fmt.Println(p4.name, p4.age, p4.Address.city, p4.Address.state)
}

func structDemo4() {
	fmt.Println("This is structDemo4---结构体指针")
	type Person struct {
		name string
		age  int
	}

	// 初始化结构体（创建一个结构体对象）
	p1 := Person{"hewuxin", 18}
	fmt.Println(p1.name, p1.age)

	// 初始化结构体指针
	// var p2 *Person  =&Person{"Hewuxin", 18}
	p2 := &Person{"heuwxin", 18}
	fmt.Println(p2.name, p2.age)

	var p3 *Person = new(Person)
	p3.name = "hewuxin"
	p3.age = 28

	fmt.Println(p3.name, p3.age)
}

func structDemo5() {
	fmt.Println("This is structDemo5---内存管理")
	type Person struct {
		name string
		age  int
	}
	// 初始化结构体
	p1 := Person{"hewuxin", 18}
	fmt.Println(p1.name, p1.age)

	// 初始化结构体指针
	p2 := &Person{"hewxin", 18}
	fmt.Println(p2.name, p2.age)
}

func structDemo6() {
	fmt.Println("This is structDemo6---赋值拷贝")
	type Person struct {
		name string
		age  int
	}
	p1 := Person{name: "hewuxin", age: 18}
	p2 := p1
	fmt.Println(p1, &p1)
	fmt.Println(p2, &p2)

	p1.name = "heyuyang"
	fmt.Println(p1, &p1)
	fmt.Println(p2, &p2)
}

func structDemo7() {
	fmt.Println("This is structDemo7---结构体指针赋值")
	type Person struct {
		name string
		age  int
	}

	p1 := &Person{"heyuyang", 19}
	p2 := p1
	fmt.Println("p1 ", p1, *p1)
	fmt.Println("p2 ", p2, *p2)

	p1.age = 20
	fmt.Println("p1 ", p1, *p1)
	fmt.Println("p2 ", p2, *p2)
}

func structDemo8() {
	fmt.Println("This is structDemo8---基于结合结构体和结构体指针的特性，基于指针实现数据变化后同步变化。")
	type Person struct {
		name string
		age  int
	}

	p1 := Person{"hewuxin", 18}
	p2 := &p1
	fmt.Println("p1 ", p1)
	fmt.Println("p2 ", p2, *p2)

	p1.name = "heyuyang"
	fmt.Println("p1 ", p1)
	fmt.Println("p2 ", p2, *p2)

}

func structDemo9() {
	fmt.Println("This is structDemo9---嵌套赋值拷贝。")
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		address Address
	}

	p1 := Person{"hewuxin", 18, Address{"beijing", "china"}}
	p2 := &p1
	fmt.Println("p1 ", p1, "p1.address ", p1.address)
	fmt.Println("p2 ", p2, " *p2 ", *p2, "*p2.address", p2.address)
	p3 := *p2
	fmt.Println("p3 ", p3, p3.address)

	p1.name = "heyuyang"
	fmt.Println("p1 ", p1, "p1.address ", p1.address)
	fmt.Println("p2 ", p2, " *p2 ", *p2, "*p2.address", p2.address)
}

func structDemo10() {
	fmt.Println("This is structDemo10")
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		hobby   [2]string
		num     []int
		parent  map[string]string
		address Address
	}
	p1 := Person{
		name:   "hey",
		age:    19,
		hobby:  [2]string{"ba", "soc"},                                // 拷贝
		num:    []int{1, 2, 3, 4, 5},                                  // 未拷贝  内部维护了指针指向数据存储的地方
		parent: map[string]string{"father": "sda", "mother": "dasda"}, // 未拷贝  内部维护了指针指向数据存储的地方
	}
	p2 := p1

	fmt.Println(p1)
	fmt.Println(p2)

	p1.hobby[0] = "dsad"
	fmt.Println(p1)
	fmt.Println(p2)

	p1.num[0] = 111111
	fmt.Println(p1)
	fmt.Println(p2)

	p1.parent["father"] = "aaaaaa"
	fmt.Println(p1)
	fmt.Println(p2)

	p1.address.city = "shanghai "
	fmt.Println(p1)
	fmt.Println(p2)
}

func structDemo11() {
	fmt.Println("This is structDemo11")
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		hobby   *[2]string
		num     []int
		parent  map[string]string
		address Address
	}
	p1 := Person{
		name:   "hey",
		age:    19,
		hobby:  &[2]string{"ba", "soc"},                               // 拷贝
		num:    []int{1, 2, 3, 4, 5},                                  // 未拷贝  内部维护了指针指向数据存储的地方
		parent: map[string]string{"father": "sda", "mother": "dasda"}, // 未拷贝  内部维护了指针指向数据存储的地方
	}
	p2 := p1

	fmt.Println(p1)
	fmt.Println(p2)

	p1.hobby[0] = "dsad"
	fmt.Println(p1)
	fmt.Println(p2)
}

func structDemo12() {
	fmt.Println("This is structDemo12")
	type Person struct {
		name string "姓名"
		age  int    "年龄"
	}

	p1 := Person{"heyuyang", 19}
	plType := reflect.TypeOf(p1)
	// 方式1
	field1 := plType.Field(0)
	fmt.Println(field1.Tag)
	//方式2
	field2, _ := plType.FieldByName("name")
	fmt.Println(field2.Tag)
	// 循环获取
	fieldNum := plType.NumField()
	for index := 0; index < fieldNum; index++ {
		filed := plType.Field(index)
		fmt.Println(filed.Name, filed.Tag)
	}

}

func structDemo() {

	structDemo9()
}
