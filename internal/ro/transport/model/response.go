package model

type Response struct {
	Code int         `json:"code"`
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
