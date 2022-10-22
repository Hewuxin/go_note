package main

import (
	"fmt"
	"reflect"
)

func descir_map() {
	fmt.Println("This is Map's demo")
}

func map_demo1() {
	fmt.Println("Map 声明和初始化")
	userInfo := map[string]string{}
	fmt.Println("userInfo: ", userInfo, reflect.TypeOf(userInfo))

	userInfo1 := map[string]string{"name": "何雨阳", "age": "18"}
	fmt.Println("userInfo1 : ", userInfo1)
	fmt.Println("My name is ", userInfo1["name"], " I'm", userInfo1["age"], " years old")

	userInfo1["name"] = "hewuxin"
	userInfo1["age"] = "24"
	fmt.Println("userInfo1 : ", userInfo1)
	fmt.Println("My name is ", userInfo1["name"], " I'm", userInfo1["age"], " years old")
}

func map_demo2() {
	fmt.Println("Map 使用make声明和初始化")
	data := make(map[int]int, 2)
	fmt.Println(" data", data, reflect.TypeOf(data))
	data[1] = 2
	fmt.Println("data", data)
	fmt.Println("length of data", len(data))

	data1 := make(map[int]int)

	fmt.Println("data1: ", data1)
	fmt.Println("length of data1 ", len(data1))
}

func map_demo3() {
	fmt.Println("使用 new 声明map指针")
	data := new(map[int]int)
	fmt.Println("data: ", data)
	fmt.Println("type of data1 ", reflect.TypeOf(data))

	data1 := make(map[int]int, 2)

	data = &data1
	fmt.Println("data1: ", data1)
	fmt.Println("data: ", data)
}

func map_demo4() {
	fmt.Println("声明一个key为[2]int数组，value为float32的map")
	v1 := make(map[[2]int]float32)
	fmt.Println("v1 : ", v1, reflect.TypeOf(v1))
	key := [2]int{1, 2}
	v1[key] = 3.14
	fmt.Println("v1: ", v1)

	fmt.Println("声明一个key为[2]int, value为[3]string的map")
	v2 := make(map[[2]int][3]string)
	fmt.Println("v2 : ", v2, reflect.TypeOf(v2))
	key1 := [2]int{5, 12}
	value1 := [3]string{"name", "age", "sex"}
	v2[key1] = value1

	fmt.Println("v2: ", v2)
}

func map_demo5() {
	fmt.Println("map 常用操作")
	data := map[string]string{"name": "h", "age": "18", "sex": "male"}
	value := len(data)
	fmt.Println("length of data is ", value)

	info := make(map[string]string, 10)
	info["n1"] = "n"
	info["n2"] = "nn"

	fmt.Println("length of info ", len(info))
	fmt.Println("info ", info)
	delete(info, "n1")
	fmt.Println("deleted n1 info is ", info)

	fmt.Println("使用 for range查看data key value: ")
	for key, value := range data {
		fmt.Println(key, value)
	}

	fmt.Println("使用 for range查看data key: ")
	for key := range data {
		fmt.Println(key)
	}

	fmt.Println("使用 for range查看data value: ")
	for _, value := range data {
		fmt.Println(value)
	}

	fmt.Println("map 嵌套")
	v1 := make(map[string]int)
	fmt.Println("type of v1: ", reflect.TypeOf(v1))

	v2 := make(map[string]string)
	fmt.Println("type of v2: ", reflect.TypeOf(v2))

	// key string
	// value [2]map[string]string 两个元素的数组  数组元素为map[string]string
	v3 := make(map[string][2]map[string]string)
	fmt.Println("v3: ", v3, reflect.TypeOf(v3))

	v3["descir"] = [2]map[string]string{map[string]string{"age": "18"},
		map[string]string{"name": "heyuyang "}}
	fmt.Println("v3: ", v3)

	// v4 := make(map[map[int]int]int) 不可哈希
}

func map_demo6() {
	fmt.Println("This is map_demo06")
	v1 := map[string]string{"name": "heyuyang", "age": "18"}

	v2 := v1
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Printf("v1's 内存地址 %p, v2's 内存地址 %p \n", &v1, &v2)

	v1["name"] = "hewuxin"
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
}

func Demo3() {
	fmt.Println("This is Demo3 ")
	descir_map()
	map_demo1()
	map_demo2()
	map_demo3()
	map_demo4()
	map_demo5()
	map_demo6()
}
