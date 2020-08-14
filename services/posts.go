package services

import (
	"fmt"

	"github.com/Neulhan/neulhan.com/db"
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/prisma"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
}

func WritePost(c *gin.Context) {
	var post Post
	err := c.BindJSON(&post)
	handling.Err(err)

	user, err := prisma.Client.User.FindOne(
		db.User.ID.Equals(post.UserId),
	).Exec(prisma.Ctx)
	handling.Err(err)

	fmt.Println(user)

	postObj, err := prisma.Client.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Content.Set(post.Content),
		db.Post.Author.Link(
			db.User.ID.Equals(user.ID),
		),
		// db.Post.Author.Set(user),
	).Exec(prisma.Ctx)

	c.JSON(200, gin.H{
		"result": "success",
		"data":   postObj,
	})
}

func PostList(c *gin.Context) {
	data, err := prisma.Client.Post.FindMany(
		db.Post.Title.Contains(""),
	).With(
		db.Post.Author.Fetch(),
	).Exec(prisma.Ctx)
	handling.Err(err)
	c.JSON(200, gin.H{"result": "success", "data": data})
}
