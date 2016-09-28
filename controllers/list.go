package controllers

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

type MyList struct {
	X int
	Y int
}

func (this *TaskController) List() {
	myStruct := MyList{1, 2}
	this.Data["json"] = &myStruct
	this.ServeJSON()
}
