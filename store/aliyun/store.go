package aliyun

import (
	"fmt"
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
	opts, err := NewDefaultAliOssStore()
	if err != nil {
		panic("load default config is error")
	}
	aliStore, err := NewAliOssStore(opts)
	if err != nil {
		panic(err)
	}
	Storer = aliStore
}

type AliOssStore struct {
	client *oss.Client
}

type Options struct {
	EndPoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) validate() error {
	if o.AccessKey == "" || o.AccessSecret == "" || o.EndPoint == "" {
		return fmt.Errorf("oss connection params is empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*Options, error) {
	opts := &Options{
		EndPoint:     conf.OssEndPoint,
		AccessSecret: conf.OssAccessSecret,
		AccessKey:    conf.OssAccessKeyId,
	}
	return opts, nil
}

// NewAliOssStore AliOssStore Constructor
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// validate input params
	if err := opts.validate(); err != nil {
		return nil, err
	}
	// new oss client
	client, err := oss.New(opts.EndPoint, opts.AccessKey, opts.AccessSecret)
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
