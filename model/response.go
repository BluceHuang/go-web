package model

type GetScheduleResponse struct {
	TYPE       int      `json:"type" bson:"type"` // 0 广告 1 活动
	ID         string   `json:"id" bson:"id"`
}

type ResponseError struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
}

type Response struct {
	Seqno string `json:"seqno"`
	Cmd string `json:"cmd"`
	Status int `json:"status"`
	Msg string `json:"msg"`
	Version version `json:"version"`
	Data reqData `json:"data"`
}