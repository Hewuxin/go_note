package main

import "fmt"

func func_demo() {
	fmt.Println("This is funcDemo")
}

func add(num1 int, num2 int) (int, bool) {
	result := num1 + num2
	return result, true
}

func funcDemo1() {
	data, flag := add(1, 8)
	fmt.Println(data, flag)
}

func ChangeData(dataList *[3]string) {
	dataList[1] = "dasda"
}

func funcDemo2() {
	userList := [3]string{"heyuyang", "hewuxin", "daxinznag"}
	ChangeData(&userList)
	fmt.Println(userList)
}
