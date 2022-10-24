package main

import "fmt"

func func_demo() {
	//fmt.Println("This is funcDemo")
	//func_demo1()
	//func_demo2()
	//func_demo3()
	//func_demo4()
	//func_demo5()

	//funcDemo6()
	//funcDemo7()
	//funcDemo8()
	funcDemo9()
	funcDemo10()

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
	//Proxy 输入参数 int, int, func 返回参数int
	// 函数exec作为第三个参数 输入参数 int, int 返回参数int,bool
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
	// exec 类型起别名 f1
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

func doDemo(num ...int) int {
	fmt.Println(num)
	sum := 0
	for _, value := range num {
		sum += value
	}

	return sum
}

func funcDemo6() {
	fmt.Println("输入变长参数Demo")
	fmt.Println("注意事项：变长参数在一个函数中只能出现一次，而且当有多个参数时，变长参数只能放在最后")
	res1 := doDemo(1, 2, 3, 5)
	res2 := doDemo(0, 1, 1)

	fmt.Println(res1, res2)
}

func add101(arg int) (int, string, bool) {
	return arg + 100, "add101", true
}

func funcDemo7() {
	fmt.Println("funcDemo7 函数返回值有多个Demo")
	v1, v2, v3 := add101(100)
	if v3 {
		message := fmt.Sprintf("add101 return %s, %d", v1, v2)
		fmt.Println(message)
	}
}

func exec(num1 int, num2 int) int {
	res := num1 + num2
	fmt.Println("这是 exec函数")
	return res
}

func getFunction() func(int, int) int {
	return exec
}

func funcDemo8() {
	fmt.Println("This is funcDemo8 函数返回值为函数")
	// getFunction 返回的是函数 使用变量存储getFunction返回的函数
	function := getFunction()
	// 使用function 传入参数计算 return
	result := function(11, 22)

	fmt.Println(result)
}

func F1(num1 int, num2 int) func(int) string {
	// 传入参数 num1 num2 int
	// 返回参数 func func传入参数int 返回参数string
	return func(num1 int) string {
		fmt.Println("匿名函数")
		return "匿名函数"
	}
}

func funcDemo9() {
	fmt.Println("This is funcDemo9 匿名函数")
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(12, 13)
	fmt.Println("data: ", data)

	value := func(n1 int, n2 int) int {
		return 123
	}(11, 22)
	fmt.Println("value :", value)

	v3 := F1(11, 22)
	res := v3(5)
	fmt.Println("F1 return :", res)
}

func droxy() func() int {
	// 通过使用droxy 调用 v1函数
	v1 := func() int {
		return 100
	}
	return v1
}

func funcDemo10() {
	fmt.Println("This is funcDemo10  使用函数作为函数的返回值")
	function := droxy() // function = v1

	result := function()
	fmt.Println("result: ", result)
}

func funcDemo11() {

	// 存储着5个函数
	var functionList []func()

	for i := 0; i < 5; i++ {
		function := func() { // 函数内部的代码只有在执行函数时才执行
			fmt.Println(i)
		}
		functionList = append(functionList, function)
	}
	//运行函数前，循环已经执行完i=5
	functionList[0]() // 5
	functionList[1]() // 5
	functionList[2]() //

}

func funcDemo12() {
	var functionList []func()

	for i := 0; i < 5; i++ {

		// 定义了匿名函数:传入参数 arg int 返回值 函数  函数为匿名函数 打印arg
		// 将i拷贝一份作为arg传入匿名函数
		function := func(arg int) func() {
			return func() {
				fmt.Println(arg)

			}
		}(i)
		functionList = append(functionList, function)
	}
	functionList[0]() // 0
	functionList[1]() // 1
	functionList[2]() // 2
}
