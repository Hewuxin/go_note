package main

import (
	"fmt"
	// "reflect"
	"unsafe"
)

func main() {
	// 字典声明
	// userInfo := map[string]string{}

	// userInfo["name"] = "heyuyang"

	// fmt.Println(userInfo)
	// fmt.Println(userInfo["name"])

	// user1 := map[string]string{"name": "hewuxin", "age": "18"}
	// fmt.Println(user1)

	// num1 := make(map[int]int, 10)
	// num1[11] = 111
	// num1[22] = 222
	// fmt.Println(num1)

	// num2 := make(map[string]int)
	// fmt.Println(num2)
	// num2["math"] = 98
	// num2["english"] = 100
	// fmt.Println(num2)

	// num1 := make(map[int]int, 10)
	// num1[11] = 111
	// num1[22] = 222
	// var row map[int]int
	// fmt.Println(row)
	// row = num1
	// fmt.Println(row)

	// value := new(map[string]int)
	// fmt.Println(value)
	// socre := make(map[string]int)
	// socre["math"] = 98
	// value = &socre
	// fmt.Println(value, reflect.TypeOf(value))

	// v1 := make(map[[2]int]float32)
	// v1[[2]int{1, 1}] = 1.1
	// fmt.Println(v1)
	// fmt.Println(v1[[2]int{1, 1}])

	// v2 := make(map[[2]int][3]int)
	// fmt.Println(v2, reflect.TypeOf(v2))
	// v2[[2]int{1, 1}] = [3]int{1, 1, 1}
	// fmt.Println(v2, v2[[2]int{1, 1}], reflect.TypeOf(v2[[2]int{1, 1}]))

	// value := map[string]string{"name": "heyuyang", "age": "24"}
	// fmt.Println(len(value))

	// value := make(map[string]string, 10)
	// value["address"] = "China"
	// fmt.Println(len(value))

	// value := map[string]string{"name": "heyuyang", "age": "24"}
	// value["age"] = "18"
	// fmt.Println(value)

	// value := map[string]string{"name": "heyuyang", "age": "24", "sex": "man", "address": "China"}
	// delete(value, "sex")
	// fmt.Println(value)

	// value := map[string]string{"name": "heyuyang", "age": "24", "sex": "man", "address": "China"}
	// for key, data := range value {
	// 	fmt.Printf("key is %v and data is %v \n", key, data)
	// }

	// for key := range value {
	// 	fmt.Println(key)
	// }

	// for _, data := range value {
	// 	fmt.Println(data)
	// }

	// v7 := make(map[string][2]map[string]string)
	// // fmt.Println(v7)
	// v7["n1"] = [2]map[string]string{{"name": "heyuyang", "age": "18"},
	// 	{"name": "hewuxin", "age": "24"}}
	// v7["n2"] = [2]map[string]string{{"name": "alex", "age": "18"},
	// 	{"name": "bob", "age": "24"}}
	// // fmt.Println(v7)
	// for key, value := range v7 {
	// 	fmt.Println(key, value)
	// }

	// v1 := map[string]string{"name": "heyuyang", "age": "18"}
	// v2 := v1
	// fmt.Printf("v1 is %v\n", v1)
	// fmt.Printf("v2 is %v\n", v2)
	// v2["age"] = "20"
	// fmt.Printf("v1 is %v\n", v1)
	// fmt.Printf("v2 is %v\n", v2)
	// v1["sex"] = "male"
	// fmt.Printf("v1 is %v\n", v1)
	// fmt.Printf("v2 is %v\n", v2)

	// slice := make([]int, 1, 3)
	// fmt.Println(cap(slice), len(slice))
	// array := [3]int{1, 2, 3}
	// fmt.Println(array, len(array), cap(array))

	// var array1 [3]int
	// fmt.Println(array1, len(array1), cap(array1))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("this is ", i)
	// }
	// var v1 string
	// var v2 *string
	// fmt.Println(reflect.TypeOf(v1))
	// fmt.Println(reflect.TypeOf(v2))

	// var name string = "何雨阳"
	// var pointer *string = &name
	// fmt.Println(name, reflect.TypeOf(name))
	// fmt.Println(pointer, reflect.TypeOf(pointer))
	// v1 := "何雨阳"
	// v2 := &v1
	// fmt.Println(v1, v2, *v2)

	// v1 = "hewuxin"
	// fmt.Println(v1, v2, *v2)

	// v1 := "heyuyang"
	// v2 := v1
	// fmt.Println(v1, v2)
	// v1 = "hewuxin"
	// fmt.Println(v1, v2)

	// v1 := "heyuyang"
	// v2 := &v1
	// fmt.Println(v1, v2, *v2)
	// v1 = "hewuxin"
	// fmt.Println(v1, v2, *v2)
	// name := "何雨阳"
	// changeData(&name)
	// changeData(name) // 这里传值时本质上会将name的值拷贝一份，并赋值给data
	// fmt.Println(name)

	// var username string
	// fmt.Println("请输入用户名: ")
	// fmt.Scanf("%s", &username)
	// if username == "何雨阳" {
	// 	fmt.Println("login succeeded")
	// } else {
	// 	fmt.Println("login failed")
	// }

	// name := "何雨阳"

	// var p1 *string = &name
	// var p2 **string = &p1
	// var p3 ***string = &p2
	// fmt.Println(name, &name)
	// fmt.Println(p1, &p1)
	// fmt.Println(p2, &p2)
	// fmt.Println(p3, &p3)
	// name := "何雨阳"
	// fmt.Println(name, &name)

	// var p1 *string = &name
	// fmt.Println("This is p1")
	// fmt.Println(*p1, p1, &p1)
	// *p1 = "张三"
	// fmt.Println(*p1, p1, &p1)

	// var p2 **string = &p1
	// fmt.Println("This is p2")
	// fmt.Println(**p2, *p2, p2, &p2)
	// **p2 = "啦啦"
	// fmt.Println(**p2, *p2, p2, &p2)

	// var p3 ***string = &p2
	// fmt.Println("This is p3")
	// fmt.Println(***p3, **p3, *p3, p3, &p3)
	// ***p3 = "我丢"
	// fmt.Println(***p3, **p3, *p3, p3, &p3)

	// dataList := [3]int8{11, 22, 33}
	// fmt.Println(dataList, reflect.TypeOf(dataList), &dataList, reflect.TypeOf(&dataList))
	// fmt.Println(dataList[0], reflect.TypeOf(dataList[0]), &dataList[0], reflect.TypeOf(&dataList[0]))

	dataList := [3]int8{11, 22, 33}

	var firstDataPtr *int8 = &dataList[0] // 获取第一个元素的地址(指针)

	ptr := unsafe.Pointer(firstDataPtr) // 转换成Pointer类型

	targetAddress := uintptr(ptr) + 2 //

	newPtr := unsafe.Pointer(targetAddress)

	value := (*int8)(newPtr)

	fmt.Println("最终结果为： ", *value)
}

// func changeData(ptr *string) {
// 	*ptr = "哈哈哈"
// }
// func changeData(data string) {
// 	data = "哈哈哈"
// }
