/*
func practice01() {
	fmt.Println("This is practice_01!")
	fmt.Printf("This is %v\n", "Printf") // Print默认不换行
	fmt.Println("Hello world")
	fmt.Printf("This is %v * %v = %v\n", 1, 3, 1*3)
}

func practice02() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.378)
	fmt.Print("lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Println(" years old.")

	fmt.Println("My weight on the surface of Mars is ", 149.0*0.378, "lbs, and I would be ",
		41*365/687)
}

func main() {
	fmt.Printf("This is func %v\n ", "practice01")
	practice01()
	fmt.Printf("This is func %v \n ", "practice02")
	practice02()
	fmt.Printf("Prinf is %v", 10)

}
*/

// func main() {
// 	fmt.Printf("This is func%4vbbbbbb\n", "aaa")
// 	fmt.Printf("This is func%-4vaaaaa", "bbb")
// }
// 常量和变量
// func main() {
// 	const num = 10
// 	var age = 10
// 	fmt.Println("\b")
// 	fmt.Printf("there have %v pens ", num)
// 	fmt.Printf("I'm %v year old after ten years.\n", age+34)
// }
// How long does it take to get to Mars?
// func var_const() {
// 	const lightspeed = 299792  // 定义常量
// 	var distance = 56000000  // 定义变量

// 	fmt.Println(distance/lightspeed, "seconds")
// 	distance = 401000000
// 	fmt.Println(distance/lightspeed, "seconds")
// }

// 定义多个变量
// func vary_const_var() {
// 	const year, weight, hight = 24, 160, 170
// 	// var distance = 5600000
// 	// var speed = 100800
// 	// var (
// 	// 	distance = 5600000
// 	// 	speed    = 100800
// 	// )
// 	var distance, speed = 5600000, 108000
// 	fmt.Printf("I'm %v years old, and %v kg %vcm \n", year, weight, hight)
// 	fmt.Printf("need %v seconds\n", distance/speed)
// 	distance = distance - 100
// 	speed = speed - 100
// 	fmt.Print("need ", distance, "\n")
// 	fmt.Print("need ", speed)
// }

// 赋值运算符
// func yunsuan() {
// 	var weight = 149.0
// 	var old = 24
// 	weight = weight * 0.3783
// 	fmt.Println("Weight", weight)
// 	weight *= 0.3782
// 	fmt.Println("wegiht", weight)
// 	old += 1
// 	fmt.Println("Old", old)
// 	old++
// 	fmt.Println("Old", old)
// 	old--
// 	fmt.Println("Old", old)
// }

// func auto_increment() {
// 	var age = 24
// 	age = age + 1
// 	fmt.Println(age)
// 	age += 1
// 	fmt.Println(age)
// 	age++
// 	fmt.Println(age)
// }
// 猜数

// func rand_num() {
// 	rand.Seed(time.Now().Unix())
// 	fmt.Println("output a num between 1 and 11")
// 	var num = rand.Intn(10) + 1
// 	fmt.Println(num)

// 	num = rand.Intn(10) + 1
// 	fmt.Println(num)
// 	var num1 = rand.Intn(10) + 1
// 	fmt.Println(num1)
// 	fmt.Printf("%v + %v = %v \n", num, num1, add(num, num1))
// }

// func add(a, b int) int {
// 	return a + b
// }

// func speed(distance, time float64) float64 {
// 	//
// 	return distance / time
// }
// func main() {
// 	// vary_const_var()
// 	// yunsuan()
// 	// auto_increment()
// 	// rand_num()
// 	const distance = 56000000.0
// 	const time = 28 * 24.0
// 	fmt.Println("speed = ", speed(distance, time))
// }

// var name, address, action string
// fmt.Print("请输入名称: ")
// fmt.Scanln(&name)

// fmt.Print("请输入位置: ")
// fmt.Scanln(&address)

// fmt.Print("请输入动作: ")
// reader := bufio.NewReader(os.Stdin)
// line, _, _ := reader.ReadLine()
// action = string(line)

// result := fmt.Sprintf("我叫 %s  在%s 在 %s", name, address, action)
// fmt.Println(result)
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// fmt.Println(math.Abs(-19))                // 绝对值
	// fmt.Println(math.Floor(3.14))             // 向下取整
	// fmt.Println(math.Ceil(3.14))              // 向上取整
	// fmt.Println(math.Round(3.34780))          // 就近取整
	// fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后两位
	// fmt.Println(math.Mod(11, 3))              // 取余数 11%3
	// fmt.Println(math.Pow(2, 5))               // 计算次方 2的5次方
	// fmt.Println(math.Pow10(2))                //计算10次方 如2的10放
	// fmt.Println(math.Max(1, 2))               // 两个值，取最大值
	// fmt.Println(math.Min(1, 2))               // 两个值，取最小值
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
