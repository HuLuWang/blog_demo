package api

import (
	"github.com/HuLuWang/blog_demo/global"
	"github.com/HuLuWang/blog_demo/internal/service"
	"github.com/HuLuWang/blog_demo/pkg/app"
	"github.com/HuLuWang/blog_demo/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	
	response.ToResponse(gin.H{
		"token": token,
	})
}
