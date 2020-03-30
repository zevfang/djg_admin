package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	BaseController
}

func (c *IndexController) Prepare() {
	if beego.AppConfig.String("runmode") == "prod" {
		isLogin := c.GetUser()
		if !isLogin {
			c.Redirect("/login", 302)
			return
		}
	}
}

// @router / [get]
func (c *IndexController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
