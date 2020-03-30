package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"encoding/json"
	"github.com/astaxie/beego"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	if beego.AppConfig.String("runmode") == "prod" {
		isLogin := c.GetUser()
		if !isLogin {
			c.Redirect("/login", 302)
			return
		}
	}
}

// @router / [get]
func (c *AuthController) Get() {
	c.Layout = "layout.html"
	c.TplName = "auth.html"
}

func (c *AuthController) GetAuths() {
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

func (c *AuthController) GetAuth() {

}

func (c *AuthController) AddAuth() {

}

func (c *AuthController) UpdAuth() {
	var model models.Auth
	resBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(resBody, &model)
	if err != nil {
		c.Data["json"] = utils.TableResult{
			Code:  400,
			Msg:   "序列化出错",
			Count: 0,
			Data:  nil,
		}
		c.ServeJSON()
	}
	err = models.EditAuth(model)
	if err != nil {
		c.Data["json"] = utils.TableResult{
			Code:  400,
			Msg:   "更新失败",
			Count: 0,
			Data:  nil,
		}
		c.ServeJSON()
	}

	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "更新成功",
		Count: 0,
		Data:  nil,
	}
	c.ServeJSON()
}
