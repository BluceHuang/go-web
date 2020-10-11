package model

type reqData interface{}

type auth struct {
	Token string `json:"token"`
	UserID string `json:"userid"`
	Platform string `json:"platform"`
	Appid string `json:"appid"`
}

type version struct {
	Ver string `json:"version"`
	Charset string `json:"charset"`
}

type Request struct {
	Seqno string `json:"seqno"`
	Cmd string `json:"cmd"`
	Msg string `json:"msg"`
	Auth auth 	`json:"auth"`
	Version version `json:"version"`
	Data reqData `json:"data"`
}

type GetScheduleReq struct {
	Mid int64 `json:"mid"`
	ProvinceId int `json:"provinceId"`
	CityId int `json:"cityId"`
	AreaId int `json:"areaId"`
	Rnids []string `json:"rnids"`
}

