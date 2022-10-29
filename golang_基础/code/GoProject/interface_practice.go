package main

import (
	"fmt"
)

// IMessage 定义 发送消息接口
type IMessage interface {
	send() bool
}

// Email 定义 邮件结构体
type Email struct {
	email   string
	content string
}

//  定义邮件结构体可以使用的send方法
func (p *Email) send() bool {
	fmt.Println("发送邮件提醒：", p.email, p.content)
	return true
}

// WeChat 定义维系结构体
type WeChat struct {
	username string
	content  string
}

// 定义微信结构体对应使用的方法
func (p *WeChat) send() bool {
	fmt.Println("发送消息提醒：", p.username, p.content)
	return true
}

// distribute 发送消息函数
func distribute(objectList []IMessage) {
	// 用户注册
	for _, item := range objectList {
		result := item.send()
		fmt.Println(result)
	}
}

func interFaceDemo() {
	messageObjectList := []IMessage{
		&Email{"daxinzang@163.com", "邮件消息提醒"},
		&WeChat{"hewuxin", "微信消息提醒"},
	}
	distribute(messageObjectList)

}
