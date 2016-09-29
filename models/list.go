package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type List struct {
	Id      uint
	Title   string
	Content string
	Ctime   time.Time
	Mtime   time.Time
	Belong  uint8
	Status  uint8
	Deleted uint8
}

func init() {
	orm.RegisterModel(new(List))
}


