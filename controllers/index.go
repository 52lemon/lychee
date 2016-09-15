package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "login/models"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
    _"github.com/astaxie/beego/session"
    "crypto/md5"
    "encoding/hex"
    "strings"
)

type IndexController struct {
    beego.Controller
}

func (index *IndexController) Get() {
    sess := index.StartSession()
    username := sess.Get("username")
    fmt.Println(username)
    if username == nil || username == "" {
        index.TplName = "index.tpl"
    } else {
        index.TplName = "success.tpl"
    }

}

func (index *IndexController) Post() {
    sess := index.StartSession()
    var user models.User
    inputs := index.Input()
    user.Username = inputs.Get("username")
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(inputs.Get("pwd")))
    cipherStr := md5Ctx.Sum(nil)
    fmt.Println(inputs.Get("pwd"))
    //user.Pwd = hex.EncodeToString(cipherStr)
    fmt.Println("---%s ", hex.EncodeToString(cipherStr))
    o := orm.NewOrm()
    //err := o.QueryTable(user).Filter("username", user.Username).Filter("pwd",user.Pwd)
    var  users models.User
    err := o.QueryTable(user).Filter("username", user.Username).One(&users, "Pwd")
    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
        fmt.Println(err)
        index.TplName = "error.tpl"
    } else if err == orm.ErrMultiRows {
        fmt.Printf("Returned Multi Rows Not One")
        index.TplName = "error.tpl"
    } else {
      fmt.Println("--users -%s ", users.Pwd)
      fmt.Println("--pass in database-%s ", user.Pwd)
      fmt.Println("---now pass is %s ", hex.EncodeToString(cipherStr))
      if(strings.EqualFold(users.Pwd,  hex.EncodeToString(cipherStr))){
            fmt.Println(user.Id, user.Username)
            sess.Set("username", user.Username)
            fmt.Println("username:", sess.Get("username"))
            index.TplName = "success.tpl"
      }else{
         index.TplName = "error.tpl"
         fmt.Printf("password is not right")
      }
    }
}
