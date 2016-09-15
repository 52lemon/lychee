package main

import (
    //"fmt"
    "github.com/astaxie/beego"
    _"login/routers"
    "github.com/astaxie/beego/orm"
    _"login/models"
    _"github.com/astaxie/beego/session"
)


func init() {
}

func main() {
    beego.BConfig.WebConfig.Session.SessionOn = true
    // 开启 ORM 调试模式
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)

    beego.Run()
}
