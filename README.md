# Go Gin Web include webrtc with websocket

include gin websocket webrtc


## How to run

### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://NULL/blockchain/blob/master/docs/sql/blog.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> NULL/blockchain/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> NULL/blockchain/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> NULL/blockchain/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> NULL/blockchain/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> NULL/blockchain/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> NULL/blockchain/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> NULL/blockchain/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> NULL/blockchain/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> NULL/blockchain/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> NULL/blockchain/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> NULL/blockchain/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```
Swagger doc

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Redis