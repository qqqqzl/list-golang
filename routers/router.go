package routers

import (
	"list-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")

	beego.Router("/api/list/:id:int", &controllers.ListController{}, "get:List")

	beego.Router("/api/list/add", &controllers.ListController{}, "post:Add")

	beego.Router("/api/list/update", &controllers.ListController{}, "post:Update")
}
