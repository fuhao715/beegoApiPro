package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/ziutek/mymysql/godrv"
	// "strconv"
	"time"
)

type OpSystem struct {
	Id         int       `orm:"column(ID);pk"`
	OsCode     string    `orm:"column(OS_CODE)"`
	OsName     string    `orm:"column(OS_NAME)"`
	OsIcon     string    `orm:"column(OS_ICON)"`
	CmsCode    string    `orm:"column(CMS_CODE)"`
	TvOsCode   string    `orm:"column(TV_OS_CODE)"`
	CreateTime time.Time `orm:"column(CREATE_TIME);auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"column(UPDATE_TIME);auto_now;type(timestamp)"`
}

func (p *OpSystem) TableName() string {
	return "ps_opsystem"
}

var logger = logs.NewLogger(10000)

func init() {
	logger.SetLogger("console", "")
	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	// orm.RegisterDataBase("default", "mysql", "root:letv@2014@tcp(10.154.252.153:3306)/taps")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(OpSystem))
}

func GetLink() beedb.Model {
	db, err := sql.Open("mymysql", "root:root123@tcp(127.0.0.1:3306)/taps") // user:password@tcp(127.0.0.1:3306)/hello
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	beedb.OnDebug = true
	return orm
}

func AddOpSystem(p OpSystem) int {
	p.Id = 123 // strconv.ParseInt("123", 0, 0)
	// UserList[p.Id] = &p
	return p.Id
}

func GetOpSystem(id int) (opSystem OpSystem, err error) {

	opSystem = OpSystem{Id: id}
	logger.Debug("opSystem.Id :{%v}", opSystem.Id)
	o := orm.NewOrm()
	o.Using("default")
	// o.QueryTable("ps_opsystem")
	beedb.OnDebug = true
	err = o.Read(&opSystem)
	logger.Debug("orm.RegisterDataBase error:{%v}", opSystem.OsCode)
	logger.Error("orm.RegisterDataBase error:{%v}", err)
	return opSystem, err

	/*
			db := GetLink()
			db.SetTable("PS_OPSYSTEM")
			err = db.Where("ID=?", id).Find(&opSystem)
		return
	*/
}
