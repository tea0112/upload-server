package server

import (
	"context"
	"net/http"

	helper "github.com/tea0112/upload-server/helpers"
	"github.com/tea0112/upload-server/route"
)

const DefaultPort string = ":8080"

type Server struct {
	port string
	context context.Context
}

func NewServer(ctx context.Context) Server {
	return Server{context: ctx}
}

// SetPort E.g. :8000
func (u *Server) SetPort(port string) {
	u.port = port
}

func (u Server) Run() {
	mux := http.NewServeMux()
	port := u.port

	if helper.IsEmpty(port) {
		port = DefaultPort
	}

	s := http.Server{
		Addr:    port,
		Handler: route.NewRoute(mux, u.context),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
