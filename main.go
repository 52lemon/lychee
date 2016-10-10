package main

import (
    //"fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _"lychee/models"
    _"github.com/astaxie/beego/session"
    "lychee/controllers"
    "github.com/astaxie/beego/logs"
)

type MainController struct {
    beego.Controller
}

func isLogin(this *MainController)(out bool){
    sess := this.StartSession()
    username := sess.Get("username")
    if username == nil || username == "" {
        out = false
    } else {
       out = true
    }
    return
}

func init() {
    beego.Info("init routers start ...")

    beego.Router("/regist", &controllers.RegistController{})

    beego.Router("/login", &controllers.LoginController{})

    beego.Router("/", &controllers.IndexController{})

    beego.Info("init routers end.")

}

func main() {
    l := logs.GetLogger()
    l.Println("this is a message of http")
    //an official log.Logger with prefix ORM
    logs.GetLogger("ORM").Println("this is a message of orm")

    logs.Debug("my book is bought in the year of ", 2016)
    logs.Info("this %s cat is %v years old", "yellow", 3)
    logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
    logs.Error(1024, "is a very", "good game")
    logs.Critical("oh,crash")

    beego.BConfig.WebConfig.Session.SessionOn = true
    beego.AddFuncMap("login",isLogin)
    // 开启 ORM 调试模式
    orm.Debug = true
    // 自动建表
    //orm.RunSyncdb("default", false, true)

    beego.Run()
}
