package routers

import (
	"lychee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Info("init routers start ...")

    beego.Router("/regist", &controllers.RegistController{})

    beego.Router("/login", &controllers.LoginController{})

    beego.Router("/", &controllers.IndexController{})

    beego.Info("init routers end.")
}
