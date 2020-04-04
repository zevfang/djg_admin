package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

const DRIVER = "mysql"

type Model struct {
	ID         int       `json:"id" orm:"column(id);size(100);pk"`
	CreatedOn  time.Time `json:"created_on" orm:"column(created_on);type(datetime)"`
	ModifiedOn time.Time `json:"modified_on" orm:"column(modified_on);type(datetime)"`
	DeletedOn  time.Time `json:"deleted_on" orm:"column(deleted_on);type(datetime)"`
}

func InitMySQL(uri string, isDebug bool) {
	BindingTable(uri, isDebug,
		new(Auth),
		new(Article),
	)
}

func BindingTable(connStr string, isDebug bool, models ...interface{}) {
	orm.Debug = isDebug
	orm.RegisterDataBase("default", "mysql", connStr)
	orm.RegisterModel(models...)
	// create table
	//orm.RunSyncdb("default", false, true)
	//设置连接数，空闲连接和失效

	db, err := orm.GetDB("default")
	if err != nil {
		panic(fmt.Errorf("get db error:%s", err))
	}
	db.SetConnMaxLifetime(time.Duration(1800) * time.Second) //30分钟失效
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(30)
}
