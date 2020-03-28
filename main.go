package main

import (
	"djg_admin/models"
	_ "djg_admin/routers"
	"github.com/astaxie/beego"
)

//var globalSessions *session.Manager

func main() {
	initSession()
	models.InitMySQL(beego.AppConfig.String("mysql_connect"))
	beego.SetStaticPath("/layui", "static/layui")
	beego.SetStaticPath("/modules", "static/layui/lay/modules")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/file", "static/file")
	beego.SetStaticPath("/img", "static/img")
	beego.Run()
}

func initSession() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "djgsessionid"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 12 * 3600 //默认值是 3600 秒
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 12 * 3600

	//sessionConfig := &session.ManagerConfig{
	//	CookieName:"djgsessionid",
	//	EnableSetCookie: true,
	//	Gclifetime:3600,
	//	Maxlifetime: 3600,
	//	Secure: false,
	//	CookieLifeTime: 3600,
	//	ProviderConfig: "./tmp",
	//}
	//globalSessions, _ = session.NewManager("memory",sessionConfig)
	//go globalSessions.GC()
}
