package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
)

type AuthController struct {
	BaseController
}

// @router / [get]
func (c *AuthController) Get() {
	c.Layout = "layout.html"
	c.TplName = "auth.html"
}

func (c *AuthController) GetAuths() {
	list, err := models.GetAuths()
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

}
