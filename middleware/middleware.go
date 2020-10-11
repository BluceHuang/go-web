package middleware

import (
	"goweb/log"
	"goweb/model"

	"github.com/gin-gonic/gin"
)

func RequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.Request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Error("req parameter format error")
			c.JSON(400, gin.H{"code": 1, "msg": "req format error", "data": gin.H{}})
			return
		}
		c.Set("request", &req)
		c.Next()
	}
}

func ResponseFormat() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *model.Request

		response := &model.Response{}
		req, exsits := c.Get("request")
		if exsits {
			request = req.(*model.Request)
		}
		var result interface{}
		result, exsits = c.Get("result")
		if exsits && request != nil {
			//response.Auth = request.Auth
			response.Cmd = request.Cmd
			response.Version = request.Version
			response.Seqno = request.Seqno
			p, ok := result.(*model.ResponseError)
			if ok {
				response.Msg = p.Msg
				response.Status = p.Status
				response.Data = struct{}{}
			} else {
				response.Data = result
			}
			c.JSON(200, response)
		} else {
			// c.JSON(500, "service handle error")
		}
		c.Next()
	}
}
