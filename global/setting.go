package global

import (
	"github.com/HuLuWang/blog_demo/configs"
	"github.com/HuLuWang/blog_demo/pkg/logger"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	CONFIG configs.Server
	VP     *viper.Viper
	DB     *gorm.DB
	Logger *logger.Logger
)
