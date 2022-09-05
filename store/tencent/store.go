package tencent

type TctStore struct{}

func NewTcTStore() (*TctStore, error) {
	return &TctStore{}, nil
}

func (s *TctStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
