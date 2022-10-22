package main

import (
	"bufio"
	"fmt"
	"os"
)

type School struct {
	band, city string
}

func StructPractice01() {
	var SchoolList []School

	for {
		var band, city string
		fmt.Print("请输入band : ")
		fmt.Scanf("%s \n", &band)
		if band == "q" {
			break
		}
		fmt.Print("请输入city: ")
		fmt.Scanf("%s \n", &city)

		sch := School{band: band, city: city}
		SchoolList = append(SchoolList, sch)
	}
	for index, item := range SchoolList {
		fmt.Println(index, item)
	}

}

type Class struct {
	title  string
	count  int
	school School
}

func getInput(inputString *string) {
	stdin := bufio.NewReader(os.Stdin)
	count, err := fmt.Fscan(stdin, &inputString)
	stdin.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		fmt.Printf("count:%d\n", count)
		fmt.Println("您的输入有误，请重新输入")
	}
}

func StructPractice02() {
	//1.创建学校
	var classList []Class

	//sch := School{band: "新东方", city: "beijing"}
	//2. 循环创建班级
	for {
		var sch School
		var class Class
		fmt.Printf("请输入班级title: ")
		fmt.Scanf("%s \n", &class.title)
		if class.title == "q" {
			break
		}

		fmt.Printf("请输入班级人数: ")
		fmt.Scanf("%s \n", &class.count)

		//
		fmt.Printf("请输入学校band: ")
		fmt.Scanf("%s \n", &sch.band)

		fmt.Printf("请输入学校city: ")
		fmt.Scanf("%s \n", &sch.city)
		class.school = sch
		// 加入到班级的列表(切片)
		classList = append(classList, class)

	}
	for _, item := range classList {
		message := fmt.Sprintf("%s %s 校区， %s 班级 有 %d 个学生", item.school.band, item.school.city,
			item.title, item.count)
		fmt.Println(message)
	}
}
