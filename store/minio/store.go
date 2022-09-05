package minio

type MiIOStore struct{}

func NewMiIOStore() (*MiIOStore, error) {
	return &MiIOStore{}, nil
}

func (s *MiIOStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
