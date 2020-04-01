package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"github.com/astaxie/beego"
)

type CtrController struct {
	BaseController
}

func (c *CtrController) Prepare() {
	if beego.AppConfig.String("runmode") == "prod" {
		isLogin := c.GetUser()
		if !isLogin {
			c.Redirect("/login", 302)
			return
		}
	}
}

// @router / [get]
func (c *CtrController) Get() {
	c.Layout = "layout.html"
	c.TplName = "ctr.html"
}

func (c *AuthController) GetCtrTop() {
	username := c.GetString("username", "")
	list, err := models.GetAuths(username)
	if err != nil {
		return
	}
	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "成功",
		Count: int64(len(list)),
		Data:  &list,
	}
	c.ServeJSON()
}
