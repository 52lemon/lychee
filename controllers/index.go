package controllers

import (
    //"fmt"
    _"lychee/models"
    _"github.com/astaxie/beego/session"
)

type IndexController struct {
    BaseController
}

func (this *IndexController) Get() {
    sess := this.StartSession()
    username := sess.Get("username")
    if username == nil || username == "" {
        this.TplName = "login.tpl"
    } else {
        this.TplName = "index.tpl"
    }

}

