package main

import (
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/prisma"
	"github.com/Neulhan/neulhan.com/services"
	"github.com/gin-gonic/gin"
)

func main() {
	client := prisma.ConnectDB()
	defer func() {
		err := client.Disconnect()
		handling.Err(err)
	}()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("v1")
	{
		v1.POST("/signup", services.SignUpService)
		v1.POST("/login", services.LoginService)
		v1.GET("/users", services.UserList)

		v1.POST("/write-post", services.WritePost)
		v1.GET("/posts", services.PostList)
	}
	r.Run()
}
