package models

import "time"

const (
	DOCUMENT = "document"
)

type Document struct {
	Id      int       `orm:"pk;auto"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Name    string    `orm:"size(100)"`
}

//TODO firt todo!

func (this *Document) TableName() string {
	return DOCUMENT
}
