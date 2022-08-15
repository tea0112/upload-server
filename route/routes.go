package route

import (
	"context"
	"net/http"

	helper "github.com/tea0112/upload-server/helpers"
	"github.com/tea0112/upload-server/aws"
	"github.com/tea0112/upload-server/user"
)

func NewRoute(mux *http.ServeMux, ctx context.Context) *http.ServeMux {
	fileWriter := helper.Writing{}
	aws := aws.AwsImpl{}
	userController := user.NewUserControllerImpl(ctx, fileWriter, aws)

	mux.HandleFunc("/user/batch", userController.Upload)

	return mux
}
