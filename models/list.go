package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type List struct {
	Id      uint
	Title   string
	Content string
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"`
	Mtime   time.Time `orm:"auto_now;type(datetime)"`
	Belong  uint8     `orm:"default(0)"`
	Status  uint8     `orm:"default(0)"`
	Deleted uint8     `orm:"default(0)"`
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

func (record *List) AddRecord() (uint, error) {
	o := orm.NewOrm()

	id, err := o.Insert(record)
	if err != nil {
		log.Println(err)
		return 0, err
	} else {
		return uint(id), nil
	}
}
