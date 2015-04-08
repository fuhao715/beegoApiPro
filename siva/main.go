package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "siva/docs"
	_ "siva/routers"
)

var log = logs.NewLogger(10000)

func init() {
	log.SetLogger("console", "")
	//log.EnableFuncCallDepth(true)
	orm.Debug = true
	err := orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/taps")
	if err != nil {
		log.Error("orm.RegisterDataBase error:{%v}", err)
	}
}

func main() {
	fmt.Println("hello , 中国golang   ! ")
	fmt.Println(beego.VERSION)
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.EnableAdmin = true
	beego.Run()
}
