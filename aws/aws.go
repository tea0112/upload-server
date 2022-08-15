package aws

type Aws interface {
	UploadFile(string) (error)
}

type AwsImpl struct{}

func (a AwsImpl) UploadFile(uploadFileDir string) (err error) {
	return nil
}
