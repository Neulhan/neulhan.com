package main

import (
	"github.com/Neulhan/neulhan.com/config"
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/middleware"
	"github.com/Neulhan/neulhan.com/prisma"
	"github.com/Neulhan/neulhan.com/services"
	"github.com/Neulhan/neulhan.com/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	client := prisma.ConnectDB()
	defer func() {
		err := client.Disconnect()
		handling.Err(err)
	}()
	storage.ConnectStorage()

	r := gin.Default()
	r.Use(middleware.CORSmiddleware())

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
		v1.GET("/images", services.ImageList)
		v1.POST("/delete-post/:key", services.DeletePost)

		v1.Static("/files", "./files")
		v1.POST("/test", services.Test)
	}
	r.Run()
}
