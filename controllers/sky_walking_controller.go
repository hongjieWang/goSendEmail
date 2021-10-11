package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"goSendEmail/constant"
	"goSendEmail/models"
	"goSendEmail/service/email_service"
	"strconv"
	"time"
)

type SkyWalkingController struct {
	web.Controller
}

// Post 发送邮件
func (this *SkyWalkingController) Post() {
	var skyWalking []models.SkyWalkingModel
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &skyWalking)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
		return
	}
	for _, model := range skyWalking {
		fmt.Println(model)
		msg := buildMsg(&model)
		println(msg)
		alarmSubject, _ := web.AppConfig.String("alarmSubject")
		alarmNickName, _ := web.AppConfig.String("alarmNickName")
		strings, _ := web.AppConfig.Strings("alarmEmails")
		var emails []string
		for _, s := range strings {
			emails = append(emails, s)
		}
		err = email_service.SendGoMail(emails, alarmSubject, alarmNickName, msg)
	}
	this.Ctx.WriteString("")
}

func buildMsg(skyWalking *models.SkyWalkingModel) string {
	micro := time.UnixMilli(skyWalking.StartTime).Format(constant.Format)
	return "scopeId: " + strconv.Itoa(skyWalking.ScopeId) +
		"</br>name: " + skyWalking.Name +
		"</br>id0: " + strconv.Itoa(skyWalking.Id0) +
		"</br>id1: " + strconv.Itoa(skyWalking.Id1) +
		"</br>告警规则: " + skyWalking.RuleName +
		"</br>告警消息: " + skyWalking.AlarmMessage +
		"</br>告警时间: " + micro
}
