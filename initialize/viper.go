package initialize

import (
	"fmt"
	"github.com/HuLuWang/blog_demo/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	
	v.SetConfigName("config")
	// 允许配置多个路径
	v.AddConfigPath("configs/")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	global.VP = v
}
