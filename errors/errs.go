package errs

type Error string

func (e Error) Error() string {
	return string(e)
} 

const ErrExceedBodySize = Error("exceed body size")
const ErrUploading = Error("upload is failed")