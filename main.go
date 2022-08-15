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

	server := server.NewServer(ctx)
	server.SetPort(PORT)
	server.Run()
}
