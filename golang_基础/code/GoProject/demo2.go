package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
数组相关示例
*/

func demo2_1() {
	fmt.Println("This is function demo2_1")
}

// 1. GO中int占多少字节？
// 2.整型中有符号和无符号是什么意思？
// 3.整型可以表示的最大范围是多少？超出怎么办？
// 4. 什么是nil？
// 5. 十进制是以整型存在，其他进制以字符串形式存在？如何实现进制之间的转换？
func demo_2_5() {
	num := 10
	// 十进制转二进制
	num_2 := strconv.FormatInt(int64(num), 2)
	fmt.Println(num_2)
	// 十进制转十六进制
	num_16 := strconv.FormatInt(int64(num), 16)
	fmt.Println(num_16)

	num_10, _ := strconv.ParseInt(num_2, 2, 16)
	fmt.Println(num_2, "二进制转十进制", num_10)

}

// 6.简述以下代码的意义
// var v1 int
// var v2 *int
// var v3 := new(int)
// 7. 浮点型为什么有时候无法精确表示小数？
// 8.如何使用第三个 decimal
// 9. 简述 ascii、unicode、utf-8的关系。
// 10. 判断 GO语言中的字符串是utf-8编码的字节序列。
// 11。什么是rune？
// 12. 判断：字符串是否可变？
// 13. 列举你记得的字符串常见功能？
func stringFunction() {
	names := "heyuyang"
	// 1. 获取长度
	fmt.Println("字符串字节长度 ", len(names))
	runeLength := utf8.RuneCountInString(names)
	fmt.Println("字符长度: ", runeLength)

	// 2. 判断是否以。。。开头
	prefixRes := strings.HasPrefix(names, "h")
	fmt.Println(prefixRes)
	// 3. 判断是否以...结尾
	suffixRes := strings.HasSuffix(names, "n")
	fmt.Println("suffixRes: ", suffixRes)

}

// 14. 字符串 和“字节集合”、“rune集合”如何实现相互转换？
func demo_2_14() {
	names := "超级为敌"
	targetByte := []byte(names)
	fmt.Println("字节集合： ", targetByte)
	targetRune := []rune(names)
	fmt.Println("rune集合: ", targetRune)

	runeToStrings := string(targetRune)
	fmt.Println("rune To string: ", runeToStrings)
	byteTostring := string(targetByte)
	fmt.Println(" byte To string: ", byteTostring)
}

// 15.字符串如何实现高效的拼接？
// 16.简述数组的存储原理。
// 17.根据要求写代码
/*
	names := [3]string{"Alex", "超级无敌小", "傻儿子"}
	a.根据索引获取"傻儿子"
	b.根据索引获取"Alex"
	c.将name最后一个元素更改为"大烧饼"
*/
func demo_2_17() {
	names := [3]string{"Alex", "超级无敌小", "傻儿子"}
	fmt.Println(names[2])
	fmt.Println(names[0])
	names[2] = "大烧饼"
	fmt.Println(names)
}

func demo_2_18() {
	var netData [3][2]int
	fmt.Println(netData)
}

// 19.声明一个有3个元素的数组，元素的类型是由两个元素的数组，并在数组中初始化值.

func demo_2_19() {
	var dataList [3][2]int
	fmt.Println("声明未赋值 ", dataList)
	dataList[0] = [2]int{11, 22}
	dataList[1] = [2]int{99, 88}
	fmt.Println("已赋值 ", dataList)
}

// 20.循环如下数组，并用字符串格式化输出如下内容
/*
	[
		["Alex", "que123"],
		["eric","admin111"],
		["tony", "ppllll"]
]
	最终实现输出:
		我是Alex,我的密码是que123
*/
func demo_2_20() {
	infoList := [3][2]string{
		[2]string{"Alex", "que123"},
		[2]string{"eric", "admin111"},
		[2]string{"tony", "ppllll"},
	}
	fmt.Println("手动循环")
	for i := 0; i < len(infoList); i++ {
		fmt.Println()
		fmt.Printf("我是%s, 我的密码是%s", infoList[i][0], infoList[i][1])
	}
	fmt.Println()
	fmt.Println("for range")
	for _, item := range infoList {
		fmt.Println()
		fmt.Printf("我是%s, 我的密码是%s", item[0], item[1])
	}
}

// 21.补充代码实现用户登录
func userLogin() {
	//userList表示有三个用户，每个用户都有用户名和密码，例如：用户名是Alex，密码是que123
	userList := [3][2]string{[2]string{"alex", "que123"},
		[2]string{"eric", "admin11"},
		[2]string{"tony", "pp1111"}}
	var userName string
	fmt.Print("请输入用户名： ")
	fmt.Scan(&userName)
	var password string
	fmt.Print("请输入密码： ")
	fmt.Scan(&password)
	fmt.Println(userName, password)
	var state string = "no"
	if state == "no" {
		for _, item := range userList {
			if item == [2]string{userName, password} {
				fmt.Println("登录成功!")
				state = "yes"
				break
			}
		}
	}
	if state == "no" {
		fmt.Println("用户名或密码不正确!")
	}
}

// 定义字符串，字符串以什么形式存在于Go编译器(utf-8编码)
func code_demo() {
	name := "何雨阳"

	// 何
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// 雨
	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

	// 阳
	fmt.Println(name[6])
	fmt.Println(name[7])
	fmt.Println(name[8])
}
