package main

import (
	"fmt"
	"time"
)

func TimeTest01() {
	t := time.Now()
	fmt.Printf("t: %T \n", t)
	fmt.Printf("t: %T \n", t)

}

func TimeTest02() {
	fmt.Println("打印显示出现在得时间， 其中now为time.Time类型， Month为time.Month类型")
	now := time.Now() // 获取当前时间
	fmt.Printf("current time : %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %2d:%2d:%2d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T, %T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
}

func TimeTest03() {
	fmt.Println("打印时间戳")
	now := time.Now()
	fmt.Printf("TimeStamp type:%T, TimeStamp: %v \n", now.Unix(), now.Unix())
}

func TimeTest04() {
	fmt.Println("打印纳秒")
	now := time.Now()
	fmt.Printf("TimeStamp type:%T, TimeStamp:%v\n", now.UnixNano(), now.UnixNano())
}

func TimeTest05() {
	fmt.Println("将时间戳转换为当前时间格式，实现瞬间替换")
	timeStamp := time.Now().Unix()
	timeObj := time.Unix(timeStamp, 0) // 将时间戳转换为时间格式
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %2d:%2d:%2d\n", year, month, day, hour, minute, second)
}

func add(h, m, s, mls, msc, ns time.Duration) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls +
		time.Microsecond*msc + time.Nanosecond*ns))
}

func TimeTest06() {
	fmt.Println("加减时间 注意这里不能增加年\\月\\日，仅能增加时分秒")
	add(3, 4, 5, 6, 7, 8)
}

func TimeTest07() {
	now := time.Now()
	targetTime := now.Add(time.Hour)
	fmt.Println("Sub Time: ", targetTime.Sub(now))
}

func TimeTest08() {
	fmt.Println("时间格式化")
	now := time.Now()
	// 格式化得模板为Go的出生时间2006年1月2号15点04分
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func TimeTest09() {
	fmt.Println("加载时区")
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(loc)
	}
}

func TimeTest10() {
	fmt.Println("加载时区")
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("按照指定时区和指定格式解析字符串时间")
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/11/15 20:31:58", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
