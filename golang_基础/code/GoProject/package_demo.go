package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func strconvDemo() {
	fmt.Println("This is strconvDemo")
	v1 := strconv.Itoa(124)
	fmt.Println("int转字符串", v1)

	v2, _ := strconv.Atoi("1221")
	fmt.Println("字符串转int: ", v2)

	// 字符串转int64整型
	v3, _ := strconv.ParseInt("123", 10, 0)
	fmt.Println("字符串转int64整型： ", v3)

	v4, _ := strconv.ParseBool("121")
	fmt.Println("字符串转bool: ", v4)

	v5 := strconv.FormatBool(true)
	fmt.Println("布尔类型转字符串类型", v5)
}

type jsonPerson struct {
	Name string
	Age  int
}

func jsonMarshalDemo() {
	fmt.Println("json序列化")
	v1 := []interface{}{
		"何雨阳",
		123,
		true,
		4.13,
		map[string]interface{}{
			"name": "hewuxin",
			"age":  18,
		},
		jsonPerson{"hewuxin", 666},
	}
	fmt.Println("原始数据", v1, reflect.TypeOf(v1))

	res, _ := json.Marshal(v1)
	fmt.Println("json序列化之后的数据: ", res)
	data := string(res)
	fmt.Println(data)
}

func jsonUnMarshalDemo() {
	fmt.Println("Json 反序列化")
	v1 := []interface{}{
		"何雨阳",
		123,
		true,
		4.13,
		map[string]interface{}{
			"name": "hewuxin",
			"age":  18,
		},
		jsonPerson{"hewuxin", 666},
	}
	fmt.Println("原始数据", v1, reflect.TypeOf(v1))

	res, _ := json.Marshal(v1)
	fmt.Println("json序列化之后的数据: ", res)
	var value []interface{}
	json.Unmarshal([]byte(res), &value)
	fmt.Println("反序列之后为： ", value)
}

//结构体的字段【首字母必须大写，不然序列化时读取不到】
// 想要让序列化后的字段为小写，可以在结构体中写tag实现

type packageInfo struct {
	Title string
	Count int
}

type Address struct {
	City string `json:"city"`
	Num  int    `json:"num"`
}

type Container struct {
	Addr Address
	Inf  packageInfo
}

func packageDemo1() {
	v1 := Container{
		Address{"上海", 10010},
		packageInfo{"标题", 10010},
	}

	// 1.序列化
	res, _ := json.Marshal(v1) // 序列化获取的是[]byte字节切片，需要转换为字符串.
	result := string(res)

	fmt.Println("序列化的结果 result: ", result)

	// 2. 反序列化
	content := `["Addr":{"city":"上海", "num": 10010}, "Inf":{"Title": "标题", "Count": 10010}]`

	var dataObject Container
	json.Unmarshal([]byte(content), &dataObject)
	fmt.Println(dataObject)
	fmt.Println(dataObject.Addr)
	fmt.Println(dataObject.Inf)
}

func packageDemo2() {
	// 1.像网络地址发送请求，并获取一段JSON格式数据。
	urlPath := "https://www.iedouyin.com/web/api/v2/aweme/iteminfo/item_id=6915066165949631752"
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)"+
		" Chrome/106.0.0.0 Safari/537.36")
	client := &http.Client{}
	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	content := string(body)
	fmt.Println("接收到的返回值内容： ", content)

	// 2. 将JSON格式数据反序列化为GO数据类型
	var responseObject map[string]interface{}
	json.Unmarshal(body, &responseObject)
	fmt.Println(responseObject)
}

func packageDemo3() {
	currentDate := time.Now()
	fmt.Println("CurrentDate: ", currentDate)

	dateString := currentDate.Format("2006_01_02_15_04_05")
	fmt.Println("格式化的时间为： ", dateString)
}

func packageDemo4() {
	// 2. 获取当前UTC时间 t2.Local()
	t2 := time.Now().UTC()

	fmt.Println("当天UTC时间 (Time类型): ", t2)

	// 3.创建一个时间，字符串类型 -> Time类型
	t3, _ := time.Parse("2006-01-02", "2022-10-29")
	fmt.Println("根据字符串转换为时间(Time)类型 :", t3)

	// 4. 创建一个时间
	t4 := time.Date(2022, 10, 29, 21, 47, 11, 11, time.Local)
	fmt.Println("根据字符串转换为时间（Time类型）: ", t4)

	t5 := time.Date(2022, 10, 29, 21, 47, 11, 11, time.UTC)
	fmt.Println("根据字符串转换为时间（Time类型）: ", t5)

	// 5. 时间格式化，Time类型 -> 字符串类型
	fmt.Println("格式化之后的时间为（string）类型： ", t2.Format("2006-01-02 15:04:05.000000 -0700 MST"))

	// 6. 时间增加
	t6 := t2.Add(time.Hour * 1)
	fmt.Println("t2 is : ", t2)
	fmt.Println("当前时间加+1小时(time类型)", t6)

	// 7. 时间减小
	t7 := t2.Add(-time.Minute * 1)
	fmt.Println("当前时间-1分钟(time类型): ", t7)

	// 8.时间间隔
	t8 := t2.Sub(t4)
	fmt.Println("时间间隔为（Duratio类型）:", t8)

	fmt.Println("时间间隔小时: ", t8.Hours())
	fmt.Println("时间间隔分钟: ", t8.Minutes())
	fmt.Println("时间间隔秒: ", t8.Seconds())
}

func packageDemo5() {
	fmt.Println("packageDemo6 基于os.Args可以获取传入的所有参数: ")
	fmt.Println(os.Args)
}

func packageDemo6() {
	fmt.Println("packageDemo6 从命令行获取参数 ")
	host := flag.String("h", "127.0.0.1", "主机名")
	port := flag.Int("p", 8080, "端口")

	var email string
	flag.StringVar(&email, "e", "", "邮箱")

	flag.Parse()
	fmt.Println(*port, *host, email)
}

func packageDemo7() {
	fmt.Println("正则表达式的知识:")
	fmt.Println("Go中正则相关的包：regexp的功能。")

	// 1. 根据字符串匹配
	m2, _ := regexp.MatchString("foo.*", "seafood")
	fmt.Println("根据字符串匹配: ", m2)
}

func packageDemo8() {
	// 创建单级目录，已存在则err有错
	err := os.Mkdir("packDemo", 0755)
	fmt.Println(err)

	// 创建多级目录，已存在报错
	err1 := os.MkdirAll("packDemo1/demo1", 0755)
	fmt.Println(err1)

	// 创建单级目录方法2
	if err := os.Mkdir("x2", 0755); err == nil {
		fmt.Println("文件夹创建成功")
	} else {
		fmt.Println(err)
	}

	// 创建多级目录方法2
	if err := os.MkdirAll("packDemo2/demo2", 0755); err == nil {
		fmt.Println("多级文件夹创建成功")
	} else {
		fmt.Println(err)
	}

}

func packageDemo9() {
	//删除文件和文件夹
	fmt.Println("删除文件和文件夹")
	if err := os.Remove("dele.go"); err == nil {
		fmt.Println("删除成功")
	} else {
		fmt.Println(err)
	}

	if err := os.Remove("packDemo"); err == nil {
		fmt.Println("单级目录删除成功")
	} else {
		fmt.Println(err)
	}

	if err := os.RemoveAll("packDemo2/demo2"); err == nil {
		fmt.Println("多级目录删除成功")
	} else {
		fmt.Println(err)
	}
}

func packageDemo10() {
	fmt.Println("判断文件或文件夹是否存在")
	_, err := os.Stat("packDemo2/demo2")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件或文件夹不存在")
		}
	} else {
		fmt.Println("文件或文件夹存在")
	}
}

func packageDemo11() {
	fmt.Println("路径拼接")
	filePath := path.Join("v1", "v2", "v3", "v4", "v6.exe")
	filePath1 := path.Join("v1", "v2", "v3/v4", "v6.exe")

	fmt.Println("filePath: ", filePath)
	fmt.Println("filePath1: ", filePath1)
}