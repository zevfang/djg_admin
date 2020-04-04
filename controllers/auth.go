package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
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

func (c *AuthController) GetAuthsSelect() {
	resultData := []models.AuthSelect{}
	list, err := models.GetAuthsAll()
	for _, item := range list {
		resultData = append(resultData, models.AuthSelect{
			Name:     item.Username,
			Value:    item.ID,
			Velected: "",
			Disabled: "",
		})
	}
	if err != nil {
		return
	}
	c.Data["json"] = utils.TableResult{
		Code:  0,
		Msg:   "success",
		Count: int64(len(resultData)),
		Data:  &resultData,
	}
	c.ServeJSON()
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

func (c *AuthController) AddAuth() {
	if c.Ctx.Request.Method == "POST" {
		var model models.Auth
		resBody := c.Ctx.Input.RequestBody
		err := json.Unmarshal(resBody, &model)
		model.CreatedOn = time.Now()
		if err != nil {
			c.Data["json"] = utils.TableResult{
				Code:  400,
				Msg:   "序列化出错",
				Count: 0,
				Data:  err,
			}
			c.ServeJSON()
		}
		if err := models.AddAuth(model); err != nil {
			c.Data["json"] = utils.TableResult{
				Code:  400,
				Msg:   "新增失败",
				Count: 0,
				Data:  err,
			}
			c.ServeJSON()
		}
		c.Data["json"] = utils.TableResult{
			Code:  200,
			Msg:   "新增成功",
			Count: 0,
			Data:  nil,
		}
		c.ServeJSON()
		return
	}
	c.TplName = "add_auth.html"
}

func (c *AuthController) GetAuth() {

}
