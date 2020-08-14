package prisma

import (
	"context"

	"github.com/Neulhan/neulhan.com/db"
	"github.com/Neulhan/neulhan.com/handling"
)

var Client *db.PrismaClient
var err error
var Ctx context.Context

func ConnectDB() *db.PrismaClient {
	Client = db.NewClient()
	err = Client.Connect()
	handling.Err(err)

	Ctx = context.Background()

	return Client
}
