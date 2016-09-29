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

func (*List) GetAllByUserId(userId uint) (*[]*List, int64, error) {
	o := orm.NewOrm()

	var lists []*List

	num, err := o.QueryTable("list").Filter("belong", userId).RelatedSel().All(&lists)

	if err != nil {
		return nil, 0, err
	} else {
		return &lists, num, nil
	}
}


