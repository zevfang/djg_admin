package main

import (
	"djg_admin/models"
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
	"time"
)

type SafeDict struct {
	data map[string]*StateModel
	*sync.RWMutex
}

func NewSafeDict(data map[string]*StateModel) *SafeDict {
	return &SafeDict{data, &sync.RWMutex{}}
}

func (d *SafeDict) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.data)
}

func (d *SafeDict) Put(key string, value *StateModel) (*StateModel, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	d.data[key] = value
	return old_value, ok
}

func (d *SafeDict) Get(key string) (*StateModel, bool) {
	d.RLock()
	defer d.RUnlock()
	old_value, ok := d.data[key]
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (*StateModel, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return old_value, ok
}

type StateModel struct {
	ArticleId      string
	Title          string
	UpdatedAt      string
	Status         string
	RecommendCount float64
	ViewCount      float64
	CommentCount   float64
	AhareCount     float64
	CollectCount   float64
	LikesCount     float64
	Ctr            float64
}

func main() {
	data := NewSafeDict(map[string]*StateModel{})
	auths := []models.Auth{
		{AppToken: "582efcd541b2d89e18a4a19515fbf349", AppID: "1654986395416859"},
		{AppToken: "a590aee1f3c30f8320bbf13eaa8588dd", AppID: "1654989824874273"},
	}
	ch := make(chan *StateModel)
	wg := sync.WaitGroup{}
	for _, item := range auths {
		wg.Add(1)
		go GetCTR(item.AppToken, item.AppID, ch, &wg)
	}
	go func() {
		for {
			r := <-ch
			data.Put(r.ArticleId, r)
		}
	}()
	wg.Wait()

	//输出最后结果
	b, err := json.Marshal(data.data)
	fmt.Println(string(b), err)
	//close(ch)
}

func GetCTR(app_token, app_id string, ch chan<- *StateModel, wg *sync.WaitGroup) {
	defer wg.Done()
	url := "https://baijiahao.baidu.com/builderinner/open/resource/query/articleListall"
	urlState := "http://baijiahao.baidu.com/builderinner/open/resource/query/articleStatistics"
	param := req.Param{
		"app_token":    app_token,
		"app_id":       app_id,
		"start_time":   "2020-01-01",
		"end_time":     time.Now().Format("2006-01-02"),
		"page_no":      1,
		"page_size":    5,
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
				"app_token":  app_token,
				"app_id":     app_id,
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
			model := StateModel{
				ArticleId:      v.Get("article_id").String(),
				Title:          v.Get("title").String(),
				UpdatedAt:      v.Get("updated_at").String(),
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
