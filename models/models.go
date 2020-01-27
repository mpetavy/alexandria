package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mpetavy/common"
)

const (
	DB_DRIVER = "db_driver"
	DB_CONN   = "db_connection"
	DB_IDLE   = "db_idle"
	DB_MAX    = "db_max"
)

var (
	Orm orm.Ormer
)

func InitDB(init bool, prefill bool) {
	orm.Debug = *common.FlagLogVerbose
	orm.DefaultTimeLoc = time.Local

	orm.RegisterModel(new(Contrep), new(Document), new(Component))
	orm.RegisterDataBase("default", beego.AppConfig.String(DB_DRIVER), beego.AppConfig.String(DB_CONN))

	err := orm.RunSyncdb("default", init, true)
	common.Fatal(err)

	Orm = orm.NewOrm()
	Orm.Using("default")

	if prefill {
		contrep := &Contrep{Name: "MP", Desc: "Description of MP content repository", IsOnline: true, IsCertChecked: true}
		_, err := Orm.Insert(contrep)
		common.Fatal(err)

		var documents []*Document

		for i := 0; i < 10; i++ {
			document := &Document{Name: "Document #" + strconv.Itoa(i)}
			_, err = Orm.Insert(document)
			common.Fatal(err)

			documents = append(documents, document)

			for j := 0; j < 10; j++ {
				component := &Component{Name: "Componet #" + strconv.Itoa(j) + " of Document " + document.Name}
				component.Document = document

				_, err = Orm.Insert(component)
				common.Fatal(err)
			}
		}
	}
}

func NewQueryBuilder() (orm.QueryBuilder, error) {
	result, err := orm.NewQueryBuilder(beego.AppConfig.String(DB_DRIVER))

	return result, err
}
