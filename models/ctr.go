package models

import "sync"

// SafeMap is the struct that wraps a hashmap with a RWMutex
type SafeMap struct {
	mu sync.RWMutex
	D  map[string]interface{}
}

// New creates a new Safemap instance
func NewSafeMap() *SafeMap {
	m := &SafeMap{
		D: map[string]interface{}{},
	}
	return m
}

// Set a value into the hashmap.
func (m *SafeMap) Set(key string, value interface{}) {
	m.mu.Lock()
	m.D[key] = value
	m.mu.Unlock()
}

// Get retrieves a value from the hashmap
func (m *SafeMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	val, ok := m.D[key]
	m.mu.RUnlock()
	return val, ok
}

// Has checks if hashmap has a key
func (m *SafeMap) Has(key string) bool {
	m.mu.RLock()
	_, ok := m.D[key]
	m.mu.RUnlock()
	return ok
}

// Count gets an integer of the number of keys in the hashmap
func (m *SafeMap) Count() int {
	m.mu.RLock()
	count := len(m.D)
	m.mu.RUnlock()
	return count
}

// Delete removes a key from the hashmap
func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	delete(m.D, key)
	m.mu.Unlock()
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
