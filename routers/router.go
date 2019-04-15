package routers

import (
	"github.com/astaxie/beego"
	"github.com/mpetavy/alexandria/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/archivelink", controllers.NewArchivelink())
}
