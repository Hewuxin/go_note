package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
字符串相关示例
*/

func bool_demo() {
	fmt.Println("This is function bool_demo")

	v1, err := strconv.ParseBool("t")
	fmt.Print(v1, err)

	v2 := strconv.FormatBool(true)
	fmt.Println("true format bool is ", v2)
}

func real_string() {
	var name string = "何雨阳"
	fmt.Println(name)
	fmt.Println("name[0] is ", name[0])
	fmt.Println("name[0] 的 2进制数 是 ", strconv.FormatInt(int64(name[0]), 2))
	fmt.Println("name[3] is ", name[3])
	fmt.Println("name[3] 的 2进制数 是", strconv.FormatInt(int64(name[3]), 2))
	fmt.Println("name[6] is ", name[6])
	fmt.Println("name[6] 的 2进制数 是", strconv.FormatInt(int64(name[6]), 2))

}

func get_string_length() {
	name := "何雨阳"
	fmt.Println(name, len(name))
}

func stringToSet() {
	name := "何雨阳"
	byteset := []byte(name)
	fmt.Println(name, " 转换为字符集合： ")
	fmt.Println("字符集合： ", byteset)
}

func setToString() {
	byteList := []byte{228, 189, 149, 233, 155, 168, 233, 152, 179}
	fmt.Println("字符集合为: ", byteList)
	fmt.Println("字符串为: ", string(byteList))
}

func stringToRune() {
	num := 65
	targetNum := strconv.FormatInt(int64((num)), 2)
	fmt.Println(num, targetNum)

	var name string = "heyuyang"
	runeSet := []rune(name)
	fmt.Println(name, runeSet)

	fmt.Println(runeSet[0], strconv.FormatInt(int64((runeSet[0])), 16))
	fmt.Println(runeSet[1], strconv.FormatInt(int64((runeSet[1])), 16))
	fmt.Println(runeSet[2], strconv.FormatInt(int64((runeSet[2])), 16))

	var name1 string = "何雨阳"
	runeSet1 := []rune(name1)
	fmt.Println(name, runeSet)

	fmt.Println(runeSet1[0], strconv.FormatInt(int64((runeSet1[0])), 16))
	fmt.Println(runeSet1[1], strconv.FormatInt(int64((runeSet1[1])), 16))
	fmt.Println(runeSet1[2], strconv.FormatInt(int64((runeSet1[2])), 16))
}

func runeToString() {
	runeList := []rune{104, 101, 121, 117, 121, 97, 110, 103}

	fmt.Println("runeList is ", runeList)

	targetString := string(runeList)
	fmt.Println("targetString is :", targetString)

}

func lengthOfRune() {
	name := "何雨阳"
	runeList := []rune(name)
	runeLength := utf8.RuneCountInString(name)
	fmt.Println(runeList, runeLength)
}

func stringMerge() {
	fmt.Println("string Merge: ")
	var builder strings.Builder
	builder.WriteString("hahaahaha")
	builder.WriteString("dadasdasda")
	value := builder.String()
	fmt.Println(value)
}

func stringIndex() {
	var name string = "呀哈哎嘿"
	v1 := name[0]
	fmt.Println("name[0]: ", v1)

	// 2. 切片获取字节区间
	v2 := name[1:3]
	fmt.Println("name[1:3]: ", v2)

	// 3. 手动循环获取所有字节
	fmt.Println("手动循环获取所有字节")
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	// for range 获取所有字节
	fmt.Println("for range 获取所有字节")
	for index, item := range name {
		fmt.Println(index, item)
	}

	runeSet := []rune{20309, 38632, 38451}
	fmt.Println("for range 获取runeList所有元素")
	for index, item := range runeSet {
		fmt.Println(index, item)
	}
}
