package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type List struct {
	Id      uint
	Title   string
	Content string    `orm:"default(\"\")"`
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"`
	Mtime   time.Time `orm:"auto_now;type(datetime)"`
	Belong  uint8     `orm:"default(0)"`
	Status  uint8     `orm:"default(0)"`
	Deleted uint8     `orm:"default(0)"`
}

type ListRes struct {
	Id      uint
	Title   string
	Content string
	Ctime   string
	Mtime   string
}

func init() {
	orm.RegisterModel(new(List))
}

func (*List) GetAllByUserId(userId uint) (*[]*ListRes, int64, error) {
	o := orm.NewOrm()

	var lists []*List

	num, err := o.QueryTable("list").Filter("belong", userId).OrderBy("-Ctime").RelatedSel().All(&lists)

	if err != nil {
		return nil, 0, err
	} else {
		listRes := make([]*ListRes, len(lists))
		for i, item := range lists {
			listRes[i] = &ListRes{
				item.Id,
				item.Title,
				item.Content,
				item.Ctime.Format("2006-01-02 15:04:05"),
				item.Mtime.Format("2006-01-02 15:04:05"),
			}
		}
		return &listRes, num, nil
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

func (record *List) UpdateRecord() (int64, error) {
	o := orm.NewOrm()

	oldList := List{Id:(*record).Id}
	err := o.Read(&oldList)
	if (err == nil) {
		oldList.Title = (*record).Title
		oldList.Content = (*record).Content
		if effectRows, err := o.Update(&oldList); err == nil {
			return effectRows, nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}






















