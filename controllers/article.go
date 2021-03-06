package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Prepare() {
	if beego.AppConfig.String("runmode") == "prod" {
		isLogin := c.GetUser()
		if !isLogin {
			c.Redirect("/login", 302)
			return
		}
	}
}

// @router / [get]
func (c *ArticleController) Get() {
	c.TplName = "article.html"
}

func (c *ArticleController) GetArticles() {
	auth_id, _ := c.GetInt64("auth_id", -1)
	state, _ := c.GetInt64("state", -1)
	verify_on := c.GetString("verify_on", "")
	if verify_on == "审核时间" {
		verify_on = ""
	}
	page, _ := c.GetInt64("page", 1)
	limit, _ := c.GetInt64("limit", 10)
	offset := (page - 1) * limit
	articles, total, err := models.GetArticles(limit, offset, auth_id, state, verify_on)
	if err != nil {
		c.Data["json"] = utils.TableResult{
			Code:  400,
			Msg:   "参数错误",
			Count: 0,
			Data:  nil,
		}
		c.ServeJSON()
	}
	resultData := []models.FindArticle{}
	for _, item := range articles {
		var model = models.FindArticle{}
		model = item
		switch item.State {
		case utils.ARTICLE_PENDING_REVIEW:
			model.StateVal = utils.ARTICLE_PENDING_REVIEW_VAL
		case utils.ARTICLE_PUBLISHED:
			model.StateVal = utils.ARTICLE_PUBLISHED_VAL
		}
		resultData = append(resultData, model)
	}

	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "成功",
		Count: total,
		Data:  &resultData,
	}
	c.ServeJSON()
}

func (c *ArticleController) GetArticle() {

}

func (c *ArticleController) UpdArticle() {
	var model models.Article
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

	//自动填写article_id
	if len(model.BJHArticleUrl) > 0 {
		sps := strings.Split(model.BJHArticleUrl, "=")
		if len(sps) == 2 {
			model.BJHArticleId = sps[1]
		}
	} else {
		model.BJHArticleId = ""
	}

	err = models.EditArticle(model)
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

func (c *ArticleController) AddArticle() {
	var model models.Article
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

	//自动填写article_id
	if len(model.BJHArticleUrl) > 0 {
		sps := strings.Split(model.BJHArticleUrl, "=")
		if len(sps) == 2 {
			model.BJHArticleId = sps[1]
		}
	} else {
		model.BJHArticleId = ""
	}

	model.CreatedOn = time.Now()

	err = models.AddArticle(model)
	if err != nil {
		c.Data["json"] = utils.TableResult{
			Code:  400,
			Msg:   "添加失败",
			Count: 0,
			Data:  nil,
		}
		c.ServeJSON()
	}

	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "添加成功",
		Count: 0,
		Data:  nil,
	}
	c.ServeJSON()
}
