package services

import (
	"fmt"
	"mime/multipart"

	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/storage"
	"github.com/gin-gonic/gin"
)

type TF struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

func Test(c *gin.Context) {
	var tf TF
	err := c.ShouldBind(&tf)
	handling.Err(err)
	for _, file := range tf.Files {
		url, err := storage.UploadFile("images", file)
		handling.Err(err)
		fmt.Println(url)
	}

	c.JSON(200, gin.H{"result": "success"})
}
