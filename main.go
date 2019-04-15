package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/mpetavy/alexandria/models"
	_ "github.com/mpetavy/alexandria/routers"
)

func main() {
	models.InitDB(true, true)

	beego.SetLogFuncCall(true)
	beego.Run()
}
