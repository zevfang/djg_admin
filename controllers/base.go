package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetUser() bool {
	is_login := c.GetSession("is_login")
	if is_login == nil {
		return false
	}
	return true
}
