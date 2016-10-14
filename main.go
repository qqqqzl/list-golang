package main

import (
	_ "list-go/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/list?charset=utf8&loc=Asia%2FShanghai")
}

func main() {
	beego.Run()
}



