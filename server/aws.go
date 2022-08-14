package server

import (
	"bytes"
	"context"
	"log"
	"os"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewAWSS3(region, bucket string) *session.Session {
	session, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func UploadFile(ctx context.Context, uploadFileDir string) error {
	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
    var fileSize int64 = upFileInfo.Size()
    fileBuffer := make([]byte, fileSize)
    upFile.Read(fileBuffer)

	session, ok := ctx.Value("s3_session").(*session.Session)
	if !ok {
		panic("s3 session doesn't initialize")
	}
	
	s3_bucket, ok := ctx.Value("s3_session").(string)
	if !ok {
		panic("s3 bucket doesn't assign")
	}
    
    _, err = s3.New(session).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(s3_bucket),
        Key:                  aws.String(uploadFileDir),
        ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(fileBuffer),
        ContentLength:        aws.Int64(fileSize),
        ContentType:          aws.String(http.DetectContentType(fileBuffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
    })
    return err
}
