package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tea0112/upload-server/aws"
	errs "github.com/tea0112/upload-server/errors"
	helper "github.com/tea0112/upload-server/helpers"
)

const maxByteRead = 1300

type UserController interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	context    context.Context
	fileWriter helper.FileWriter
	aws        aws.Aws
}

func NewUserControllerImpl(ctx context.Context, fileWriter helper.FileWriter, aws aws.Aws) UserController {
	return UserControllerImpl{context: ctx, fileWriter: fileWriter, aws: aws}
}

func (u UserControllerImpl) Upload(w http.ResponseWriter, r *http.Request) {

	buffer := make([]byte, maxByteRead)
	counter := 0

	for {
		n, err := r.Body.Read(buffer)
		counter += n
		if counter > maxByteRead {
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

	jsonBuffer := buffer[:counter]
	var jsonBody map[string]interface{}

	ch := make(chan error)
	go func() {
		fileDir := "/tmp/file"
		err := u.fileWriter.WriteFile(fileDir, jsonBuffer)
		if err != nil {
			ch <- err
		}

		err = u.aws.UploadFile(fileDir)
		ch <- err
	}()

	err := <-ch
	if err != nil {
		http.Error(w, errs.ErrUploading.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.Unmarshal(jsonBuffer, &jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(jsonBody)

}
