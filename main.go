package main

import (
	"fmt"
	"github.com/HuLuWang/blog_demo/global"
	"github.com/HuLuWang/blog_demo/initialize"
	"log"
	"net/http"
	"time"
)

func init() {
	// init viper must be first
	initialize.Viper()
	// init log
	initialize.Logger()
	// init mysql
	initialize.Mysql()
}

func main() {
	address := fmt.Sprintf(":%d", global.CONFIG.System.HttpPort)
	log.Printf("%d", global.CONFIG.System.HttpPort)
	Router := initialize.Routers()
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    global.CONFIG.System.ReadTimeout,
		WriteTimeout:   global.CONFIG.System.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`默认自动化文档地址:http://127.0.0.1%s/swagger/index.html`, s.Addr)
	s.ListenAndServe()
}
