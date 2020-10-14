# blog_demo
go gin
```
.
├── README.md
├── configs
│   ├── config.go
│   └── config.yaml
├── docs
│   ├── docs.go
│   ├── sql
│   │   └── blog.sql
│   ├── swagger.json
│   └── swagger.yaml
├── global
│   └── setting.go
├── go.mod
├── go.sum
├── initialize
│   ├── logger.go
│   ├── mysql.go
│   ├── router.go
│   └── viper.go
├── internal
│   ├── dao
│   │   ├── article.go
│   │   ├── article_tag.go
│   │   ├── dao.go
│   │   └── tag.go
│   ├── middleware
│   │   ├── access_log.go
│   │   ├── context_timeout.go
│   │   ├── jwt.go
│   │   └── limiter.go
│   ├── model
│   │   ├── article.go
│   │   ├── article_tag.go
│   │   ├── model.go
│   │   └── tag.go
│   ├── routers
│   │   └── api
│   │       ├── auth.go
│   │       └── v1
│   │           ├── article.go
│   │           └── tag.go
│   └── service
│       ├── article.go
│       ├── auth.go
│       ├── service.go
│       └── tag.go
├── main.go
├── pkg
│   ├── app
│   │   ├── app.go
│   │   ├── form.go
│   │   ├── jwt.go
│   │   └── pagination.go
│   ├── convert
│   │   └── convert.go
│   ├── errcode
│   │   ├── common_code.go
│   │   ├── errcode.go
│   │   └── module_code.go
│   ├── limiter
│   │   ├── limiter.go
│   │   └── method_limiter.go
│   ├── logger
│   │   └── logger.go
│   └── util
│       └── md5.go
└── scripts

```

### Todo 
- 链路追踪
