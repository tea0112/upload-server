package user

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

type WritingSuccessfulMock struct{}

func (w WritingSuccessfulMock) WriteFile(path string, buffer []byte) error {
	return nil
}

type AwsSuccessfulUploadingMock struct{}

func (a AwsSuccessfulUploadingMock) UploadFile(dir string) error {
	return nil
}

func TestUpload(t *testing.T) {
	t.Run("upload success", func(t *testing.T) {
		dirname, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		f, err := os.Open(path.Join(dirname, "../payload.json"))
		if err != nil {
			panic(err)
		}

		r := httptest.NewRequest(http.MethodPost, "/", bufio.NewReader(f))
		w := httptest.NewRecorder()

		userController := UserControllerImpl{fileWriter: WritingSuccessfulMock{}, aws: AwsSuccessfulUploadingMock{}}
		userController.Upload(w, r)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("upload file is failed")
		}
	})
}
