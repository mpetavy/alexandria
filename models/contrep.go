package models

import "time"

const (
	CONTREP = "contrep"
)

type Contrep struct {
	Id                   int       `orm:"pk;auto"`
	Created              time.Time `orm:"auto_now_add;type(datetime)"`
	Updated              time.Time `orm:"auto_now;type(datetime)"`
	Name                 string    `orm:"size(100)"`
	Desc                 string    `orm:"null;size(1000)"`
	IsOnline             bool      `orm:"size(1000)"`
	Cert                 string    `orm:"null;size(2000)"`
	CertDatetime         time.Time `orm:"null"`
	IsCertProtected      bool
	IsCertChecked        bool
	ReceivedCert         string    `orm:"null;size(2000)"`
	ReceivedCertDatetime time.Time `orm:"null"`
	DocProtection        string    `orm:"null;size(4)"`
}

func (this *Contrep) TableName() string {
	return CONTREP
}
