package models

import (
    _ "database/sql"
    _"errors"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
)

type User struct {
    Id       int `PK`
    Username string
    Pwd      string
}

func init(){
    //注册 model
    orm.RegisterModel(new(User))
    //注册驱动
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:root@/login?charset=utf8")
}


