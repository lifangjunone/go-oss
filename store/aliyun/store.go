package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-oss/conf"
	"go-oss/store"
)

var (
	// Whether the object implements the constraints of the interface
	_      store.Storer = &AliOssStore{}
	Storer store.Storer
)

func init() {
	aliStore, err := NewAliOssStore(conf.OssEndPoint, conf.OssAccessKeyId, conf.OssAccessSecret)
	if err != nil {
		panic(err)
	}
	Storer = aliStore
}

type AliOssStore struct {
	client *oss.Client
}

// NewAliOssStore AliOssStore Constructor
func NewAliOssStore(endPoint, accessKey, accessSecret string) (*AliOssStore, error) {
	client, err := oss.New(endPoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client}, nil
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(objectKey, fileName)
}
