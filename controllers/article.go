package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"fmt"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Prepare() {
	isLogin := c.GetUser()
	if !isLogin {
		c.Redirect("/login", 302)
		return
	}
}

// @router / [get]
func (c *ArticleController) Get() {
	c.Layout = "layout.html"
	c.TplName = "article.html"
}

func (c *ArticleController) GetArticles() {
	page, _ := c.GetInt64("page", 1)
	limit, _ := c.GetInt64("limit", 10)
	offset := (page - 1) * limit
	articles, total, err := models.GetArticles(limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "成功",
		Count: total,
		Data:  &articles,
	}
	c.ServeJSON()
}
func (c *ArticleController) GetArticle() {

}
func (c *ArticleController) UpdArticle() {

}
