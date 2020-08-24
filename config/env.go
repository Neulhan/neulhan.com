package config

import (
	"os"

	"github.com/Neulhan/neulhan.com/handling"
	"github.com/joho/godotenv"
)

var Get map[string]string

func LoadConfig() {
	err := godotenv.Load()
	handling.Err(err)

	Get = map[string]string{
		"S3_BUCKET":         os.Getenv("S3_BUCKET"),
		"ACCESS_KEY_ID":     os.Getenv("ACCESS_KEY_ID"),
		"SECRET_ACCESS_KEY": os.Getenv("SECRET_ACCESS_KEY"),
	}
}
