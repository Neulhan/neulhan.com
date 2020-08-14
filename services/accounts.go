package services

import (
	"github.com/Neulhan/neulhan.com/db"
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/Neulhan/neulhan.com/prisma"
	"github.com/gin-gonic/gin"
)

type SignUp struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func SignUpService(c *gin.Context) {
	var data SignUp
	err := c.BindJSON(&data)
	handling.Err(err)

	createdUser, err := prisma.Client.User.CreateOne(
		db.User.Email.Set(data.Email),
		db.User.Age.Set(data.Age),
		db.User.Name.Set(data.Name),
		db.User.Password.Set(data.Password),
	).Exec(prisma.Ctx)
	handling.Err(err)

	c.JSON(200, gin.H{
		"data":        data,
		"createdUser": createdUser,
	})
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginService(c *gin.Context) {
	var login Login
	result := false
	err := c.BindJSON(&login)
	handling.Err(err)

	user, err := prisma.Client.User.FindOne(db.User.Email.Equals(login.Email)).Exec(prisma.Ctx)
	handling.Err(err)

	pw, ok := user.Password()
	if ok && pw == login.Password {
		result = true
	}

	c.JSON(200, gin.H{
		"login": result,
	})
}

func UserList(c *gin.Context) {
	users, err := prisma.Client.User.FindMany(
		db.User.Email.Contains(""),
	).With(
		db.User.Posts.Fetch(),
	).Exec(prisma.Ctx)
	handling.Err(err)
	c.JSON(200, gin.H{
		"users": users,
	})
}
