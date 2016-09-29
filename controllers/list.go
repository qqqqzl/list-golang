package controllers

import (
	"github.com/astaxie/beego"
	"list-go/models"
)

type ListController struct {
	beego.Controller
}

type Res struct {
	Errno  uint
	Errmsg string
	Data   *ResData
}

type ResData struct {
	Lists *[]*models.List
	Total int64
}

func (this *ListController) List() {

	var list models.List
	var res Res

	lists, num, err := list.GetAllByUserId(1)

	if err != nil {
		res = Res{Errno:1, Errmsg:"can not find record"}
	} else {
		resData := ResData{lists, num}
		res = Res{Errno:0, Errmsg:"success", Data:&resData}
	}
	this.Data["json"] = res
	this.ServeJSON()
}
