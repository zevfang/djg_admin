package controllers

import (
	"djg_admin/models"
	"djg_admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
	"time"
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

func (c *CtrController) GetCtrTop() {
	user_ids := c.GetString("user_ids", "")
	//updated_at := c.GetString("updated_at", "")
	page_no, _ := c.GetInt64("page_no", 1)
	page_size := int64(5) //默认只能传入5条，否则百家号api查询有误
	if len(user_ids) == 0 {
		c.Data["json"] = utils.TableResult{
			Code:  200,
			Msg:   "无条件传入",
			Count: 0,
			Data:  []models.CtrModel{},
		}
		c.ServeJSON()
		return
	}
	auths, _ := models.GetAuthByIds(user_ids)
	safeData := models.NewSafeMap()

	ch := make(chan *models.CtrModel)
	wg := sync.WaitGroup{}
	wg_count := len(auths) * int(page_size)
	wg.Add(wg_count)

	for _, item := range auths {
		go getCTR(page_no, page_size, item.Username, item.AppToken, item.AppID, ch)
	}
	go func(wg *sync.WaitGroup) {
		for {
			r := <-ch
			safeData.Set(r.ArticleId, r)
			wg.Done()
		}
	}(&wg)
	wg.Wait()

	resultData := []*models.CtrModel{}
	for _, v := range safeData.D {
		resultData = append(resultData, v.(*models.CtrModel))
	}
	c.Data["json"] = utils.TableResult{
		Code:  200,
		Msg:   "成功",
		Count: int64(len(resultData)),
		Data:  resultData,
	}
	c.ServeJSON()
}

func getCTR(pageNo, pageSize int64, userName string, appToken, appId string, ch chan<- *models.CtrModel) {
	url := "https://baijiahao.baidu.com/builderinner/open/resource/query/articleListall"
	urlState := "http://baijiahao.baidu.com/builderinner/open/resource/query/articleStatistics"
	param := req.Param{
		"app_token":    appToken,
		"app_id":       appId,
		"start_time":   "2020-01-01",
		"end_time":     time.Now().Format("2006-01-02"),
		"page_no":      pageNo,
		"page_size":    pageSize,
		"article_type": "video",
		"collection":   "publish",
	}
	lr, err := req.Get(url, param)
	if err != nil {
		fmt.Println(err)
	}
	listJson := lr.String()
	if gjson.Valid(listJson) {
		items := gjson.Get(listJson, "data").Get("items").Map()
		for _, v := range items {
			sr, err := req.Get(urlState, req.QueryParam{
				"app_token":  appToken,
				"app_id":     appId,
				"article_id": v.Get("article_id"),
			})
			if err != nil {
				fmt.Println(err)
			}
			stateJson := sr.String()
			stateJsonData := gjson.Get(stateJson, "data")
			view_count := stateJsonData.Get("view_count").Float()
			recommend_count := stateJsonData.Get("recommend_count").Float()
			ctrStr := fmt.Sprintf("%.2f", (view_count / recommend_count * 100.00))
			ctr, _ := strconv.ParseFloat(ctrStr, 64)
			//时间戳转换
			updateTime := time.Unix(v.Get("updated_at").Int(), 0).Format("2006-01-02 15:04:05")
			model := models.CtrModel{
				UserName:       userName,
				PreviewUrl:     fmt.Sprintf("%s%s", "http://baijiahao.baidu.com/builder/preview/s?id=", v.Get("article_id").String()),
				ArticleId:      v.Get("article_id").String(),
				Title:          v.Get("title").String(),
				UpdatedAt:      updateTime,
				Status:         v.Get("status").String(),
				RecommendCount: recommend_count,
				ViewCount:      view_count,
				CommentCount:   stateJsonData.Get("comment_count").Float(),
				AhareCount:     stateJsonData.Get("share_count").Float(),
				CollectCount:   stateJsonData.Get("collect_count").Float(),
				LikesCount:     stateJsonData.Get("likes_count").Float(),
				Ctr:            ctr,
			}
			ch <- &model
		}
	}
}
