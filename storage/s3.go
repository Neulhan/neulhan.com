package storage

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/Neulhan/neulhan.com/config"
	"github.com/Neulhan/neulhan.com/handling"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3Endpint = ""

var s3Client *session.Session
var err error

func ConnectStorage() {
	cred := credentials.NewStaticCredentials(config.Get["ACCESS_KEY_ID"], config.Get["SECRET_ACCESS_KEY"], "")
	fmt.Println(cred)
	s3Client, err = session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.ApNortheast2RegionID),
		Credentials: cred,
	})
	handling.Err(err)
}

func UploadFile(path string, fileheader *multipart.FileHeader) (string, error) {
	file, err := fileheader.Open()
	handling.Err(err)
	defer file.Close()
	var size = fileheader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := t.Month()
	day := strconv.Itoa(t.Day())
	wholePath := path + "/" + year + "/" + month.String() + "/" + day + "/" + fileheader.Filename
	_, s3err := s3.New(s3Client).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.Get["S3_BUCKET"]),
		Key:    aws.String(wholePath),
		// ACL:           aws.String("public"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		// ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: aws.String("AES256"),
		// StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	S3EndPoint := config.Get["S3_ENDPOINT"]
	return S3EndPoint + wholePath, s3err
}
