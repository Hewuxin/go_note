package main

import "fmt"

// IMessage 创建send接口
type IMessage interface {
	send() bool
}

type Email struct {
	email   string
	content string
}

func (p *Email) send() bool {
	fmt.Println("发送邮件提醒：", p.email, p.content)
	return true
}

type WeChat struct {
	userName string
	content  string
}

func (p *WeChat) send() bool {
	fmt.Println("发送wechat提醒: ", p.userName, p.content)
	return true
}

func distribute(objectList []IMessage) {
	// 用户注册
	for _, item := range objectList {
		result := item.send()
		fmt.Println(result)
	}
}

func messageDis() {
	email := Email{email: "py_daxinzang@163.com", content: "发送邮件"}
	wechat := WeChat{userName: "打心脏", content: "发送微信"}

}
