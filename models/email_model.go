package models

// EmailModel 发送短信对象
type EmailModel struct {
	//接收地址集合
	MailAddress []string
	//主题
	Subject string
	//消息
	Body string
	//别名
	Nickname string
}
