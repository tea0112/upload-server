package main

import (
	"context"

	"github.com/tea0112/upload-server/server"
)


const (
	AWS_S3_REGION = ""
	AWS_S3_BUCKET = ""
	PORT = ":8080"
)

func main() {
	ctx := context.Background()

	session := server.NewAWSS3(AWS_S3_REGION, AWS_S3_BUCKET)
	ctx = context.WithValue(ctx, "s3_session", session)
	ctx = context.WithValue(ctx, "s3_bucket", AWS_S3_BUCKET)

	server := server.NewServer(ctx)
	server.SetPort(PORT)
	server.Run()
}
