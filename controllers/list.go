package controllers

import (
	"github.com/astaxie/beego"
	"list-go/models"
	"strconv"
	"github.com/astaxie/beego/validation"
	"log"
)

type ListController struct {
	beego.Controller
}

type Res struct {
	Errno  uint
	Errmsg string
	Data   interface{}
}

type ResData struct {
	Lists *[]*models.ListRes
	Total int64
}

func (this *ListController) List() {

	userId, err := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 0)

	if err != nil {
		res := Res{Errno:1, Errmsg:"invalid user id"}
		this.Data["json"] = res
		this.ServeJSON()
	} else {
		lists, num, err := (&models.List{}).GetAllByUserId(uint(userId))
		var res Res
		if err != nil {
			res = Res{Errno:1, Errmsg:"can not find record"}
		} else {
			resData := ResData{lists, num}
			res = Res{Errno:0, Errmsg:"success", Data:&resData}
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}

func (this *ListController) Add() {
	title := this.GetString("title")
	content := this.GetString("content")

	valid := validation.Validation{}
	valid.Required(title, "title")

	var errmsg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errmsg += " " + err.Key + " : " + err.Message + " ; "
			log.Println(err.Key, err.Message)
		}
		res := Res{Errno:1, Errmsg:errmsg}
		this.Data["json"] = res
		this.ServeJSON()
	} else {
		record := models.List{
			Title:title,
			Content:content,
		}

		recordId, err := record.AddRecord()
		var res Res
		if err != nil {
			res = Res{Errno:1, Errmsg:"add record fail"}
		} else {
			res = Res{Errno:0, Errmsg:"success", Data:struct {
				Id uint
			}{recordId}}
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}





















