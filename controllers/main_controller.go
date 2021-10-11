package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (main *MainController) Get() {
	main.Ctx.WriteString("hello GeeGo!")
}

