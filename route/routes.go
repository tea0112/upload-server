package route

import (
	"context"
	"net/http"
	"github.com/tea0112/upload-server/user"
)

func NewRoute(mux *http.ServeMux, ctx context.Context) *http.ServeMux {
	userController := user.NewUserControllerImpl(ctx)

	mux.HandleFunc("/user/batch", userController.Upload)

	return mux
}
