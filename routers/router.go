package routers

import (
	"list-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/list", &controllers.ListController{}, "get:List")
}
