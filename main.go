package main

import (
    //"fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _"lychee/models"
    _"github.com/astaxie/beego/session"
    "lychee/controllers"
)


func init() {
    beego.Info("init routers start ...")

    beego.Router("/regist", &controllers.RegistController{})

    beego.Router("/login", &controllers.LoginController{})

    beego.Router("/", &controllers.IndexController{})

    beego.Info("init routers end.")

}

func main() {
    beego.BConfig.WebConfig.Session.SessionOn = true
    // 开启 ORM 调试模式
    orm.Debug = true
    // 自动建表
    //orm.RunSyncdb("default", false, true)

    beego.Run()
}
