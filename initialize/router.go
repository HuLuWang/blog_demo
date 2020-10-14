package initialize

import (
	_ "github.com/HuLuWang/blog_demo/docs"
	"github.com/HuLuWang/blog_demo/global"
	"github.com/HuLuWang/blog_demo/internal/middleware"
	"github.com/HuLuWang/blog_demo/internal/routers/api"
	v1 "github.com/HuLuWang/blog_demo/internal/routers/api/v1"
	"github.com/HuLuWang/blog_demo/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// 初始化限流器
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     100,
		Quantum:      100,
	},
)

//初始化总路由
func Routers() *gin.Engine {
	var r = gin.New()
	if global.CONFIG.System.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(gin.Recovery())
	}
	// 限流器
	r.Use(middleware.RateLimiter(methodLimiters))
	// 超时控制
	r.Use(middleware.ContextTimeout(60))
	tag := v1.NewTag()
	article := v1.NewArticle()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)
	
	apiv1 := r.Group("/api/v1")
	// 接入jwt
	//apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/tags", tag.List)         //标签列表
		apiv1.POST("/tags", tag.Create)       //创建标签
		apiv1.PUT("/tag/:id", tag.Update)    //更新标签
		apiv1.DELETE("/tag/:id", tag.Delete) //删除标签
		
		apiv1.GET("/articles", article.List)         //文章列表
		apiv1.POST("/articles", article.Create)       //创建文章
		apiv1.PUT("/article/:id", article.Update)    //更新文章
		apiv1.DELETE("/article/:id", article.Delete) //删除文章
	}
	return r
}
