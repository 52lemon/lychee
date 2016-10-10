package controllers

import (
    "fmt"
    "lychee/models"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
    _"github.com/astaxie/beego/session"
    "crypto/md5"
    "encoding/hex"
    "strings"
)

type LoginController struct {
    BaseController
}

func (this *LoginController) Get() {
    this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
    sess := this.StartSession()
    var user models.User
    inputs := this.Input()
    user.Username = inputs.Get("username")
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(inputs.Get("pwd")))
    cipherStr := md5Ctx.Sum(nil)
    fmt.Println(inputs.Get("pwd"))
    fmt.Println("---%s ", hex.EncodeToString(cipherStr))
    o := orm.NewOrm()
    var  users models.User
    err := o.QueryTable(user).Filter("username", user.Username).One(&users, "Pwd")
    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
        fmt.Println(err)
        this.TplName = "error.tpl"
    } else if err == orm.ErrMultiRows {
        fmt.Printf("Returned Multi Rows Not One")
        this.TplName = "error.tpl"
    } else {
      fmt.Println("--users -%s ", users.Pwd)
      fmt.Println("--pass in database-%s ", user.Pwd)
      fmt.Println("---now pass is %s ", hex.EncodeToString(cipherStr))
      if(strings.EqualFold(users.Pwd,  hex.EncodeToString(cipherStr))){
            fmt.Println(user.Id, user.Username)
            sess.Set("username", user.Username)
            fmt.Println("username:", sess.Get("username"))
            this.TplName = "index.tpl"
      }else{
         this.TplName = "error.tpl"
         fmt.Printf("password is not right")
      }
    }
}
