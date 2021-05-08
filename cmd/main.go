package main

import (
	"net/http"
	"short_url/service"
	"short_url/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	index := router.Group("")
	{
		shortService := service.ShortService{}
		index.POST("/add", func(ctx *gin.Context) {
			url := ctx.Query("url")
			ok, redirect := shortService.CreateShortURL(url)
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"msg": "程序错误"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"msg": "操作成功", "data": redirect})
		})
		index.GET("/:code", func(ctx *gin.Context) {
			code := ctx.Param("code")
			url := shortService.RedirectURL(code)
			if url == "" {
				ctx.JSON(http.StatusOK, gin.H{"msg": "程序错误"})
				return
			}
			ctx.Redirect(http.StatusMovedPermanently, url)
		})
	}
	// 初始化redis
	tool.RedisInit()
	// 检查错误
	err := router.Run(":9998")
	if err != nil {
		return
	}
}
