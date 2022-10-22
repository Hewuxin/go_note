package main

import "fmt"

func func_demo() {
	fmt.Println("This is funcDemo")
}

func add(num1 int, num2 int) (int, bool) {
	result := num1 + num2
	return result, true
}

func func_demo1() {
	fmt.Println("输入两个参数 返回两个参数")
	data, flag := add(1, 2)
	fmt.Println(data, flag)
}

func SendEmail(arg [2]int) {
	arg[0] = 5
}

func func_demo2() {
	fmt.Println("传值时会拷贝一份数据 等于赋值拷贝")
	dataList := [2]int{1, 3}
	fmt.Println(dataList)
	SendEmail(dataList)
	fmt.Println(dataList)

}

func SendEmailPointer(arg *[2]int) {
	arg[0] = 10
}

func func_demo3() {
	fmt.Println("---------------------指针参数-------------------------")
	fmt.Println("要是通过函数改变变量的值需要传入变量的内存地址")
	dataList := [2]int{1, 3}
	fmt.Println("dataList: ", dataList)
	SendEmailPointer(&dataList)
	fmt.Println("dataList: ", dataList)

}

func add100(num1 int, num2 int) (int, bool) {
	res := num1 + num2
	return res, true
}

func proxy(num1 int, num2 int, exec func(num1 int, num2 int) (int, bool)) int {
	data, flag := exec(num1, num2)
	if flag {
		return data
	} else {
		return data - 100
	}
}
func func_demo4() {
	fmt.Println("函数做参数传入函数")
	res := proxy(100, 200, add100)
	fmt.Println(res)
}

type f1 func(num1 int, num2 int) (int, bool)

func sroxy(num1 int, num2 int, exec f1) int {
	data, flag := exec(num1, num2)
	if flag {
		return data
	} else {
		return 1
	}
}

func func_demo5() {
	fmt.Println("函数做参数传入函数")
	res := sroxy(1, 50, add100)
	fmt.Println(res)
}
