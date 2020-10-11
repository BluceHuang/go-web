package controller

import (
	"goweb/log"
	"goweb/model"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) interface{} {
	var request *model.Request
	//_ = ctx.ShouldBindJSON(&req)
	req, exists := ctx.Get("request")
	if !exists {
		log.Error("req parameter not exists")
		return nil
	}
	request = req.(*model.Request)
	return request.Data
}

func returnError(ctx *gin.Context, status int, msg string) {
	repError := &model.ResponseError{Status: status, Msg: msg}
	ctx.Set("result", repError)
}

func returnResult(ctx *gin.Context, result interface{}) {
	ctx.Set("result", result)
}
