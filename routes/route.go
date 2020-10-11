package routes

import (
	"goweb/middleware"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func BuildRoute() *gin.Engine {
	return engine
}

func GetRouter(relationPath string) *gin.RouterGroup {
	return engine.Group(relationPath)
}

func init() {
	engine = gin.Default()
	engine.Use(middleware.RequestBody())
}
