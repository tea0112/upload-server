package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tea0112/upload-server/errors"
)

const maxByteRead = 1146

type UserController interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	context context.Context
}

func NewUserControllerImpl(ctx context.Context) UserController {
	return UserControllerImpl{context: ctx}
}

func (u UserControllerImpl) Upload(w http.ResponseWriter, r *http.Request) {
	buffer := make([]byte, maxByteRead)
	counter := 0	

	for {
		n, err := r.Body.Read(buffer)
		counter += n
		if counter >= maxByteRead {
			http.Error(w, errs.ErrExceedBodySize.Error(), http.StatusBadRequest)
			return
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	var jsonBody map[string]interface{}

	err := json.Unmarshal(buffer[:counter], &jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonBody)
	w.WriteHeader(http.StatusOK)
}
