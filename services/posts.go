package services

import (
	"mime/multipart"

	"github.com/Neulhan/neulhan.com/db"
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/prisma"
	"github.com/Neulhan/neulhan.com/storage"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Title   string                  `form:"title" binding:"required"`
	Content string                  `form:"content" binding:"required"`
	UserId  string                  `form:"userId" binding:"required"`
	Images  []*multipart.FileHeader `form:"images" binding:"required"`
}

func WritePost(c *gin.Context) {
	var post Post
	err := c.ShouldBind(&post)
	handling.Err(err)

	user, err := prisma.Client.User.FindOne(
		db.User.ID.Equals(post.UserId),
	).Exec(prisma.Ctx)
	handling.Err(err)

	postObj, err := prisma.Client.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Content.Set(post.Content),
		db.Post.Author.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(prisma.Ctx)
	handling.Err(err)

	for _, file := range post.Images {
		url, err := storage.UploadFile("images", file)
		handling.Err(err)

		_, err = prisma.Client.Image.CreateOne(
			db.Image.Name.Set(file.Filename),
			db.Image.URL.Set(url),
			db.Image.Post.Link(
				db.Post.ID.Equals(postObj.ID),
			),
		).Exec(prisma.Ctx)
		handling.Err(err)
	}

	postObj, err = prisma.Client.Post.FindOne(
		db.Post.ID.Equals(postObj.ID),
	).With(
		db.Post.Author.Fetch(),
		db.Post.Images.Fetch(),
	).Exec(prisma.Ctx)

	c.JSON(200, gin.H{
		"result": "success",
		"data":   postObj,
	})
}
func DeletePost(c *gin.Context) {
	key := c.Param("key")
	_, err := prisma.Client.Image.FindMany(
		db.Image.PostID.Equals(key),
	).Delete().Exec(prisma.Ctx)
	handling.Err(err)
	_, err = prisma.Client.Post.FindOne(
		db.Post.ID.Equals(key),
	).Delete().Exec(prisma.Ctx)
	handling.Err(err)

	c.JSON(200, gin.H{
		"result": "success",
	})
}

func PostList(c *gin.Context) {
	data, err := prisma.Client.Post.FindMany(
		db.Post.Title.Contains(""),
	).With(
		db.Post.Author.Fetch(),
		db.Post.Images.Fetch(),
	).Exec(prisma.Ctx)
	handling.Err(err)
	c.JSON(200, gin.H{"result": "success", "data": data})
}

func ImageList(c *gin.Context) {
	data, err := prisma.Client.Image.FindMany(
		db.Image.Name.Contains(""),
	).Exec(prisma.Ctx)
	handling.Err(err)
	c.JSON(200, gin.H{"result": "success", "data": data})
}
