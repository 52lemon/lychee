package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "login/models"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
    "crypto/md5"
    "encoding/hex"
)

type RegistController struct {
    beego.Controller
}

func (this *RegistController) Get() {
    this.TplName = "regist.tpl"
}

func (this *RegistController) Post() {
    var user models.User
    inputs := this.Input()
    //fmt.Println(inputs)
    user.Username = inputs.Get("username")
    //var pass = inputs.Get("pwd")
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(inputs.Get("pwd")))
    cipherStr := md5Ctx.Sum(nil)
    fmt.Println(cipherStr)
    fmt.Println(hex.EncodeToString(cipherStr))
    user.Pwd = hex.EncodeToString(cipherStr)
    o := orm.NewOrm() 
    fmt.Println(user)
    _, err := o.Insert(&user)
    if err == nil {
        this.TplName = "success.tpl"
    } else {
        fmt.Println(err)
        this.TplName = "error.tpl"
    }
}
