package email_service

import (
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/gomail.v2"
)

// MailHost 邮件服务器地址
var MailHost, _ = web.AppConfig.String("mailHost")

// MailPort 端口
var MailPort, _ = web.AppConfig.Int("mailPort")

// MailUser 发送邮件用户账号
var MailUser, _ = web.AppConfig.String("mailUser")

// MailPwd 授权密码
var MailPwd, _ = web.AppConfig.String("mailPwd")

// SendGoMail /*
func SendGoMail(mailAddress []string, subject string, nickname string, body string) error {
	m := gomail.NewMessage()

	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	m.SetHeader("From", nickname+"<"+MailUser+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)

	d := gomail.NewDialer(MailHost, MailPort, MailUser, MailPwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
