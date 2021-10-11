package routers

import (
	"github.com/beego/beego/v2/server/web"
	"goSendEmail/controllers"
)

func Init() {
	web.Router("/hello", &controllers.MainController{})
	web.Router("/sendEmail", &controllers.EmailController{})
	web.Router("/alarm", &controllers.SkyWalkingController{})
}
