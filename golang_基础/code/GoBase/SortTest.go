package main

import (
	"fmt"
	"sort"
)

/*

 */

func SortTest() {
	s := []int{2, 4, 1, 5}
	sort.Ints(s)
	fmt.Printf("s: %v \n", s)
}

type NewInts []uint // 定义一个新接口 切片类型 元素为uint

// Len 为NewInts绑定函数
func (n NewInts) Len() int {
	return len(n)
}

//Less 为NewInts绑定函数
func (n NewInts) Less(i, j int) bool {
	fmt.Println("定义比较函数")
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

// Swap 为NewInts绑定函数
func (n NewInts) Swap(i, j int) {
	fmt.Println("定义交换函数")
	n[i], n[j] = n[j], n[i]
}

func SortTest01() {
	n := []uint{1, 3, 2}
	sort.Sort(NewInts(n))
	fmt.Println(n)
}

func SortTest02() {
	fmt.Println("元素为float64的切片排序")
	f := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	sort.Float64s(f)
	fmt.Printf("f : %v \n", f)
}

func SortTest03() {
	fmt.Println("元素为int的切片排序")
	f := []int{1, 4, 5, 3, 2}
	sort.Ints(f)
	fmt.Printf("f : %v \n", f)
}

func SortTest04() {
	fmt.Println("字符串排序，先比较高位，相同再比较低位")
	ls := sort.StringSlice{
		"100",
		"42",
		"3",
		"-1",
	}
	fmt.Println(ls)
	sort.Strings(ls)
	fmt.Printf("ls : %v \n", ls)
}

func SortTest05() {
	fmt.Println("字符串排序，先比较高位，相同再比较低位")
	ls := sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls)
	sort.Strings(ls)
	fmt.Printf("ls : %v \n", ls)
}

func SortTest06() {
	fmt.Println("汉字排序，依次比较byte大小")
	ls := sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
	}
	fmt.Println(ls)
	sort.Strings(ls)
	fmt.Printf("ls : %v \n", ls)
}

type testSlice [][]int

func (l testSlice) Len() int {
	return len(l)
}

func (l testSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l testSlice) Less(i, j int) bool {
	return l[i][1] < l[j][1]
}

func SortTest07() {
	fmt.Println("复杂结构排序")
	ls := testSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	}

	fmt.Println(ls)
	sort.Sort(ls)
	fmt.Println(ls)
}