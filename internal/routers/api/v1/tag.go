package v1

import (
	"github.com/HuLuWang/blog_demo/global"
	"github.com/HuLuWang/blog_demo/pkg/app"
	"github.com/HuLuWang/blog_demo/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 创建
// @Tags Tag 标签
// @Accept application/json
// @Produce application/json
// @Param name body string true "标签" maxlength(20)
// @Param created_by body string false "创建者"
// @Success 200
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 标签List
// @Tags Tag 标签
// @Accept application/json
// @Produce application/json
// @Param page query int false "页数"
// @Param page_size query int false "条数"
// @Success 200
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新
// @Tags Tag 标签
// @Accept application/json
// @Produce application/json
// @Param id path int true "标签ID"
// @Param name body string false "标签" maxlength(20)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "修改者"
// @Success 200
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除
// @Tags Tag 标签
// @Accept application/json
// @Produce application/json
// @Param id path int true "标签ID"
// @Success 200
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
