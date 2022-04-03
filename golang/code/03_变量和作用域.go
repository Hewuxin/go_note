// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var era = "AD"

// func main() {
// 	fmt.Println("This is variable and scope")
// variable()
// short_statment()
// if_short_use()
// switch_short_use()
// date_show()
// }

// func date_show() {
// 	year := 2018
// 	rand.Seed(time.Now().Unix())
// 	switch month := rand.Intn(12) + 1; month {
// 	case 2:
// 		day := rand.Intn(28) + 1
// 		fmt.Println(era, year, month, day)
// 	case 4, 6, 9, 11:
// 		day := rand.Intn(30) + 1
// 		fmt.Println(era, year, month, day)
// 	default:
// 		day := rand.Intn(31) + 1
// 		fmt.Println(era, year, month, day)
// 	}
// }

// func switch_short_use() {
// 	rand.Seed(time.Now().Unix())
// 	var num = 2
// 	switch num = rand.Intn(6) + 1; num {
// 	case 1:
// 		fmt.Println("this is num", num)
// 	case 2:
// 		fmt.Println("This is num ", num)
// 	case 3:
// 		fmt.Println("this is num", num)
// 	case 4:
// 		fmt.Println("this is num", num)
// 	case 5:
// 		fmt.Println("this is num", num)
// 	default:
// 		fmt.Println("not have num", num)
// 	}

// }

// func if_short_use() {
// 	// var num = 0
// 	rand.Seed(time.Now().Unix())
// 	// if num = rand.Intn(3); num == 0 {
// 	// 	fmt.Println(num)
// 	// 	fmt.Println("Space Adventure")
// 	// }
// 	if num := rand.Intn(3); num == 0 {
// 		fmt.Println(num)
// 		fmt.Println("Space Adventures")
// 	}

// }

// func short_statment() {
// 	// for count := 5; count > 0; count-- {
// 	// 	fmt.Println(count)
// 	// 	time.Sleep(time.Second)
// 	// }
// 	var count = 0
// 	for count = 5; count > 0; count-- {
// 		fmt.Println(count)
// 		time.Sleep(time.Second)

// 	}
// }

// func variable() {
// 	var count = 0

// 	for count < 10 {
// 		fmt.Println(count)
// 		count++
// 		time.Sleep(time.Second)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var era = "AD"

// func main() {
// 	randomdate_show()

// }

// func randomdate_show() {
// 	for count := 10; count > 0; count-- {
// 		// year := 2018
// 		rand.Seed(time.Now().Unix())
// 		year := rand.Intn(22) + 2001
// 		month := rand.Intn(12) + 1
// 		daysInmonth := 31

// 		switch month {
// 		case 2:
// 			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
// 				fmt.Printf("%v is leap year \n", year)
// 				daysInmonth = 29
// 			} else {
// 				daysInmonth = 28
// 			}

// 		case 4, 6, 9, 11:
// 			daysInmonth = 30

// 		}
// 		day := rand.Intn(daysInmonth) + 1
// 		fmt.Println(era, year, month, day)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	name := "heyuyang"
// 	nickname := name

// 	fmt.Println(name, &name)
// 	name = "hewuxin"
// 	fmt.Println(name, &name)
// 	fmt.Println(nickname, &nickname)
// }

package main

import "fmt"

var number = 99

func main() {
	var name string
	fmt.Print("请输入姓名")
	fmt.Scanln(&name)

	if name == "wupeiqi" {
		goto SVIP
	} else if name == "yuanhao" {
		goto VIP
	}
	fmt.Println("预约。。。。。。")
VIP:
	fmt.Println("等号。。。。。。。")
SVIP:
	fmt.Println("进入..........")
}
