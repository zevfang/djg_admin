package controllers

import (
	"djg_admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 128
	cpt.StdHeight = 35
}

type LoginController struct {
	beego.Controller
}

// @router / [get]
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

// @router / [post]
func (c *LoginController) Post() {
	var remember = c.GetString("remember", "off")
	var username = c.GetString("username")
	var password = c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = &utils.Result{
			Code: 400,
			Msg:  "账号或密码为空",
		}
		c.ServeJSON()
	}
	isVerify := cpt.VerifyReq(c.Ctx.Request)
	if !isVerify {
		c.Data["json"] = &utils.Result{
			Code: 400,
			Msg:  "验证码错误",
		}
		c.ServeJSON()
	}

	if username == "root" && password == "fangwei0505" {
		//记住我
		if remember == "on" {

		}
		c.SetSession("is_login", username)
		c.Data["json"] = &utils.Result{
			Code: 200,
			Msg:  "登录成功",
		}
		c.ServeJSON()
	}

	c.Data["json"] = &utils.Result{
		Code: 400,
		Msg:  "未知错误，登录失败！",
	}
	c.ServeJSON()
}

func (c *LoginController) LoginOut() {
	c.DelSession("is_login")
	c.Redirect("/login", 302)
}
