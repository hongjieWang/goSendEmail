package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"goSendEmail/models"
	"goSendEmail/service/email_service"
)

// EmailController 发送邮件服务
type EmailController struct {
	web.Controller
}

// Post 发送邮件
func (this *EmailController) Post() {
	var email models.EmailModel
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &email)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
		return
	}
	err = email_service.SendGoMail(email.MailAddress, email.Subject, email.Nickname, email.Body)
	if err != nil {
		this.Ctx.WriteString("email send error !" + err.Error())
		return
	}
	this.Ctx.WriteString("email send success !")
}
