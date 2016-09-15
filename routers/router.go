package routers

import (
	"login/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Info("init routers start ...")

    beego.Router("/", &controllers.IndexController{})
    beego.Router("/regist", &controllers.RegistController{})

    beego.Info("init routers end.")
}
