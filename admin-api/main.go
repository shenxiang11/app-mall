package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/sx931210/app-mall/admin-api/api"
	_ "github.com/sx931210/app-mall/admin-api/docs"
)

// @title 电商后台
// @description admin-api
// @version 1.0
// @contact.name shen
// @contact.url https://github.com/shenxiang11
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))

	api.RegisterHandlers(r)

	r.Run(":8080")
}
