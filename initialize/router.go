package initialize

import "github.com/gin-gonic/gin"

//初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	return Router
}

