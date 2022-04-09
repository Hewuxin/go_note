package main

import (
	"fmt"
	"math"
)

func main() {
	// message := "shalom"
	// for count := 0; count < len(message); count++ {
	// 	fmt.Printf("%c\n", message[count])
	// }
	fmt.Println(math.Abs(-19))                // 绝对值
	fmt.Println(math.Floor(3.14))             // 向下取整
	fmt.Println(math.Ceil(3.14))              // 向上取整
	fmt.Println(math.Round(3.34780))          // 就近取整
	fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后两位
	fmt.Println(math.Mod(11, 3))              // 取余数 11%3
	fmt.Println(math.Pow(2, 5))               // 计算次方 2的5次方
	fmt.Println(math.Pow10(2))                //计算10次方 如2的10放
	fmt.Println(math.Max(1, 2))               // 两个值，取最大值
	fmt.Println(math.Min(1, 2))               // 两个值，取最小值
}
