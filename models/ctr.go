package models

import "sync"

// 安全的map
type SafeDict struct {
	Data map[string]*CtrModel
	*sync.RWMutex
}

func NewSafeDict(data map[string]*CtrModel) *SafeDict {
	return &SafeDict{data, &sync.RWMutex{}}
}

func (d *SafeDict) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.Data)
}

func (d *SafeDict) Put(key string, value *CtrModel) (*CtrModel, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.Data[key]
	d.Data[key] = value
	return old_value, ok
}

func (d *SafeDict) Get(key string) (*CtrModel, bool) {
	d.RLock()
	defer d.RUnlock()
	old_value, ok := d.Data[key]
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (*CtrModel, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.Data[key]
	if ok {
		delete(d.Data, key)
	}
	return old_value, ok
}

type CtrModel struct {
	UserName       string  `json:"user_name"`
	PreviewUrl     string  `json:"preview_url"`
	ArticleId      string  `json:"article_id"`
	Title          string  `json:"title"`
	UpdatedAt      string  `json:"updated_at"`
	Status         string  `json:"status"`
	RecommendCount float64 `json:"recommend_count"`
	ViewCount      float64 `json:"view_count"`
	CommentCount   float64 `json:"comment_count"`
	AhareCount     float64 `json:"ahare_count"`
	CollectCount   float64 `json:"collect_count"`
	LikesCount     float64 `json:"likes_count"`
	Ctr            float64 `json:"ctr"`
}
