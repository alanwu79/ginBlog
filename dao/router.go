package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})

	return r
}
