package models

import "time"

const (
	COMPONENT = "component"
)

type Component struct {
	Id       int       `orm:"pk;auto"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Name     string    `orm:"size(100)"`
	Document *Document `orm:"rel(fk)"`
}

func (this *Component) TableName() string {
	return COMPONENT
}
