package initialize

import (
	_ "github.com/HuLuWang/blog_demo/docs"
	"github.com/HuLuWang/blog_demo/internal/middleware"
	"github.com/HuLuWang/blog_demo/internal/routers/api"
	v1 "github.com/HuLuWang/blog_demo/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.POST("/auth", api.GetAuth)
	
	
	tag := v1.NewTag()
	article := v1.NewArticle()
	apiv1 := Router.Group("/api/v1")
	// 接入jwt
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/tags", tag.List)         //标签列表
		apiv1.POST("/tag", tag.Create)       //创建标签
		apiv1.PUT("/tag/:id", tag.Update)    //更新标签
		apiv1.DELETE("/tag/:id", tag.Delete) //删除标签
		
		apiv1.GET("/articles", article.List)         //文章列表
		apiv1.POST("/article", article.Create)       //创建文章
		apiv1.PUT("/article/:id", article.Update)    //更新文章
		apiv1.DELETE("/article/:id", article.Delete) //删除文章
	}
	return Router
}
