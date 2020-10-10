package service

import (
	"context"
	"github.com/HuLuWang/blog_demo/internal/dao"
	"github.com/HuLuWang/blog_demo/global"
)

//
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DB)
	return svc
}
