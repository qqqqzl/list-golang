package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"list-go/models"
)

type ListController struct {
	beego.Controller
}

type Res struct {
	Errno  uint
	Errmsg string
	Data   *models.List
}

func (this *ListController) List() {
	o := orm.NewOrm()
	o.Using("list")

	list := models.List{Id:1}
	err := o.Read(&list)
	var res Res
	if err == orm.ErrNoRows {
		res = Res{Errno:1, Errmsg:"can not find record"}
	} else if err == orm.ErrMissPK {
		res = Res{Errno:2, Errmsg:"can not find primary key"}
	} else {
		res = Res{Errno:0, Errmsg:"success", Data:&list}
	}
	this.Data["json"] = res
	this.ServeJSON()
}
