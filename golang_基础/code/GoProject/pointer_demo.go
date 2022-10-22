package main

import (
	"fmt"
	"reflect"
)

func Demo1() {
	fmt.Println("This is Demo2")
	var v1 string
	fmt.Println("v1: ", v1, reflect.TypeOf(v1))

	var v2 *string
	fmt.Println("v1: ", v2, reflect.TypeOf(v2))

	var name string = "Heyuyang"
	var pointer *string = &name
	fmt.Println("name is ", name)
	fmt.Println("pointer is ", pointer, "value is ", *pointer)

	var age int = 18
	var x1 *int = &age
	fmt.Println("age is ", age)
	fmt.Println("x1 is ", pointer, "x1's value is ", *x1)
}

func Demo2() {
	fmt.Println("This is Demo2")
	v1 := "heyuyang"
	v2 := &v1

	fmt.Println("v1 ", v1, " v2 ", v2, " *v2 ", *v2)

	v1 = "hewuxin"
	fmt.Println("v1 ", v1, " v2 ", v2, " *v2 ", *v2)
}

func pointerDemo3() {
	fmt.Println("This is pointerDemo3")
	v1 := "heyuyang"
	v2 := v1
	fmt.Println("v1: ", v1, " v2 ", v2)
	v1 = "hewuxin"
	fmt.Println("v1: ", v1, " v2 ", v2)
	fmt.Println("pointer")
	v3 := "hewuxin"
	v4 := &v3
	fmt.Println("v3: ", v3, " v4 ", v4, " *v4 ", *v4)
}

func changeData(data string) {
	fmt.Println("This is changeData")
	data = "hahahah"
	fmt.Println("changeData's data is ", data)
}

func Demo4() {
	fmt.Println("This is pointerDemo4")
	name := "何雨阳"
	changeData(name)
	fmt.Println("name is ", name)

}

func changePointerData(ptr *string) {
	fmt.Println("This is changePointerData")
	*ptr = "哈哈哈"
	fmt.Println("changeData's ptr is ", ptr, " ptr's value is ", *ptr)
}

func Demo5() {
	fmt.Println("This is pointerDemo5")
	name := "heyuyang"
	fmt.Println("Before name is ", name)
	changePointerData(&name)
	fmt.Println("After changePointerData name is ", name)
}

func Demo6() {
	name := "hewuxin"
	fmt.Println("This is pointerDemo6")
	var p1 *string = &name

	var p2 **string = &p1

	var p3 ***string = &p2

	fmt.Printf("name is %s p1 is %p p1's value is %s p2 is %p p2's value is %p p3 is %p p3's value is %p \n",
		name, p1, *p1, p2, *p2, p3, *p3)
	fmt.Println(name, &name)
	fmt.Println(p1, &p1)
	fmt.Println(p2, &p2)
	fmt.Println(p3, *p3)
}

func Demo7() {
	fmt.Println("This is pointerDemo7")

	var username string
	fmt.Println("请输入username: ")
	fmt.Scanf("%s", &username)

	if username == "hewuxin" {
		fmt.Println("login successful!")
	} else {
		fmt.Println("login failure.")
	}
}

func Demo8() {
	fmt.Println("This is pointerDemo8")
	name := "heyuyang"

	var p1 *string = &name

	*p1 = "hewuxin"
	fmt.Println("This is p1 ", p1, " *p1 ", *p1, " &p1 ", &p1)

	var p2 **string = &p1
	**p2 = "拉拉"
	fmt.Println("This is p2 ", p2, " *p2 ", *p2, " &p2 ", &p2)

	var p3 ***string = &p2
	***p3 = "wodiu"
	fmt.Println("This is p3 ", p3, " *p3 ", *p3, " &p3 ", &p3)

}

func pointerDemo() {
	fmt.Println("This is Pointer Demo")
	//Demo1()
	//Demo2()
	//pointerDemo3()
	//Demo4()
	//Demo5()
	Demo6()
	//Demo7()
	Demo8()
}
