package main

import (
	"fmt"
	"reflect"
)

func desci() {
	fmt.Println("This is slice demo")
}

func slice_new() {
	var users = []int{11, 22, 33}
	fmt.Println(users)
	fmt.Println("users's length: ", len(users))
	fmt.Println("users's cap: ", cap(users))

	data := []int{33, 44, 55}
	fmt.Println(data, len(data), cap(data))
	fmt.Println("type of data ", reflect.TypeOf(data))

	var data1 = make([]int, 1, 3)
	fmt.Println("data1 : ", data1, len(data1), cap(data1))
	var v2 = new([]int)
	fmt.Println("type of v2 ", reflect.TypeOf(v2))

	var v3 *[]int
	fmt.Println("type of v2 ", reflect.TypeOf(v3))

	// 扩容前和扩容后地址不同

	v4 := make([]int, 1, 1)
	fmt.Println("type of v4 ", reflect.TypeOf(v4))
	fmt.Printf("v4 %p \n", &v4)
	fmt.Println(v4)
	v4 = append(v4, 2)
	fmt.Println("type of v4 ", reflect.TypeOf(v4))
	fmt.Printf("v4 %p \n", &v4)
	fmt.Println(v4)

	v5 := new([]int)
	fmt.Println(v5)
	fmt.Println(*v5)
	fmt.Println("type of v5 ", reflect.TypeOf(v5))
}

func slice_3() {
	v1 := make([]int, 2, 5)
	fmt.Println("slice_3: ", v1[0], len(v1), cap(v1)) // 0, 2, 5
}

func slice_4() {
	v1 := make([]int, 2, 5)
	v2 := append(v1, 124)

	fmt.Println("slice_4: ")
	fmt.Println(len(v1), cap(v1)) // 2, 5
	fmt.Println(len(v2), cap(v2)) // 3, 5
}

func slice_5() {
	v1 := make([]int, 2, 5)
	v2 := append(v1, 124)
	v1[0] = 999
	fmt.Println("slice_5: ")
	fmt.Println(v1) //  [999 0]
	fmt.Println(v2) // [999,0,124]
}

func slice_6() {
	v1 := make([]int, 2, 2)
	v2 := append(v1, 124)
	v1[0] = 999
	fmt.Println("slice_6: ")
	fmt.Println(v1) // [999 0]
	fmt.Println(v2) // [0 0 124]
}

func slice_7() {
	v1 := make([]int, 2, 2)
	v2 := v1[0:2]
	v1[0] = 111

	fmt.Println("slice_7: ")
	fmt.Println(v1) // [111, 0]
	fmt.Println(v2) // [111, 0]
}

func slice_8() {
	v1 := [][]int{[]int{11, 22, 33}, []int{44, 55}}
	v1[0][2] = 999

	fmt.Println("slice_8: ")
	fmt.Println(v1) // [[11 22 999][44 55]]
}

func slice_9() {
	v1 := [][]int{[]int{11, 22, 33}, []int{44, 55}}
	v2 := v1[0]
	v2[2] = 69

	fmt.Println("slice_9: ")
	fmt.Println(v1) // [[11 22 69][44 55]]
}

func slice_10() {
	v1 := [][]int{make([]int, 2, 5), make([]int, 2, 5)}
	v2 := append(v1[0], 99)
	v2[2] = 69

	fmt.Println("slice_10: ")
	fmt.Println(v1) // [[0 0][0 0]]
}

func slice_11() {
	v1 := [][]int{make([]int, 2, 5), make([]int, 2, 5)}
	v2 := append(v1[0], 99)

	v2[0] = 69

	fmt.Println("slice_11: ")
	// [[0 0][0 0]]wrong  v1 和v2共用同一片内存 只是保存的len不同 [[69 0 ][0 0]]
	fmt.Println(v1)
	fmt.Println(v2) // [69 0 99]
}

func slice_func() {
	v1 := []int{11, 22, 33}

	v2 := append(v1, 44)
	v2[0] = 0
	v3 := append(v1, 55, 66, 77, 88)

	v4 := append(v1, []int{123, 45, 67, 89}...)

	fmt.Println("slice_func: ")
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}

func slice_func1() {
	v1 := make([]int, 2, 5)

	v2 := append(v1, 44)
	v2[1] = 100
	v3 := append(v1, 55, 66, 77, 88)

	v4 := append(v1, []int{123, 45, 67, 89}...)
	v2[0] = 1
	fmt.Println("slice_func1: ")
	fmt.Println(v1) // [1 100]
	fmt.Println(v2) // [1 100 44]
	fmt.Println(v3) // [0 100 55 66 77 88]
	fmt.Println(v4) // [0 100 123 45 67 89]
}
