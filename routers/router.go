package routers

import (
	"list-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/list/:id:int", &controllers.ListController{}, "get:List")

	beego.Router("/api/list/add", &controllers.ListController{}, "post:Add")
}
