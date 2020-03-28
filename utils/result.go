package utils

type TableResult struct {
	Code  int64       `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
