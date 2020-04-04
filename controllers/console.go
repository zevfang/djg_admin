package controllers

import "github.com/astaxie/beego"

type ConsoleController struct {
	BaseController
}

func (c *ConsoleController) Prepare() {
	if beego.AppConfig.String("runmode") == "prod" {
		isLogin := c.GetUser()
		if !isLogin {
			c.Redirect("/login", 302)
			return
		}
	}
}

// @router / [get]
func (c *ConsoleController) Get() {
	c.TplName = "console.html"
}
