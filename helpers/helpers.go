package helper

import "os"

func IsEmpty(str string) bool {
	if str == "" {
		return true
	} else {
		return false
	}
}

type FileWriter interface {
	WriteFile(string, []byte) error
}

type Writing struct{}

func (w Writing) WriteFile(path string, buffer []byte) error {
	f, err := os.Create("/tmp/file")
	if err != nil {
		return err
	}

	_, err = f.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
