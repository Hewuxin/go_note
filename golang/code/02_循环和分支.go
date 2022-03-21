package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func contains() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outsid"
	var exit = strings.Contains(command, "outsid")
	fmt.Println("You leave cave: ", exit)
}

func is_minor() {
	fmt.Println("There is a sign the entrance that reads 'No Minors'")
	var age = 24
	var minor = age < 18
	fmt.Printf("At age %v , am I a minor? %v \n", age, minor)

}

func is_east(command string) {
	if command == "go east" {
		fmt.Println("You head further up the mountain.")

	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")

	} else {
		fmt.Println("Didn't quite get that.")
	}
}

// 逻辑运算符
func yu() {
	var year = get_year()

	fmt.Printf("The year is %d, should you leap?\n", year)

	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap! ")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}

func get_year() int {
	var year int
	fmt.Println("please enter a year num")
	fmt.Scanln(&year)
	fmt.Println("you enterd the num is ", year)
	return year
}

// 取反运算符
func qufan() {
	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
}

// 使用switch分支
func switch_use() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
		fallthrough // 不进行判断就执行下一个case语句。
	case "read sign":
		fmt.Println("The sign reads 'No Minors.'")
	default:
		fmt.Println("Didn't quite get that.")
	}
}

func fallthrough_use() {
	var num int
	fmt.Println("please enter a num to chose do something!")
	fmt.Scanln(&num)

	switch {
	case num == 1:
		fmt.Println("play basketball")
	case num == 2:
		fmt.Println("play football")
	case num == 3:
		fmt.Println("go shopping")
		fallthrough
	case num == 4:
		fmt.Println("play baseball")
	}
}

// for 循环  后面没跟条件为无线循环  后面跟条件为有限循环

func for_use() {
	var count = 10
	for count > 0 {
		if count < 5 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second) // time.Second 等于1秒
		count--
	}
	fmt.Println("Litoff!")
}

// 猜数
func homework() {
	rand.Seed(time.Now().Unix())
	var num int
	num = rand.Intn(100) + 1
	for num > 0 {
		var scan_num int
		fmt.Print("please enter a num: ")
		fmt.Scanln(&scan_num)
		if num == scan_num {
			fmt.Println("猜对了")
			break
		} else if num > scan_num {
			fmt.Println("太小了")
		} else if num < scan_num {
			fmt.Println("太大了")
		}

	}
}
func main() {
	fmt.Println("This is func main")
	contains()
	is_minor()
	var command string
	var command1 string
	fmt.Println("please enter the command")
	fmt.Scanf("%s %s", &command, &command1)
	fmt.Println(command + command1)
	is_east(command + " " + command1)
	yu()
	qufan()
	switch_use()
	fallthrough_use()
	for_use()
	homework()
}
