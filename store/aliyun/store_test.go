package aliyun

import (
	"github.com/stretchr/testify/assert"
	"go-oss/conf"
	"testing"
)

// TestAliOssStore_Upload
func TestAliOssStore_Upload(t *testing.T) {
	// use assert
	should := assert.New(t)
	err := Storer.Upload(conf.OssBucket, "store.go", "store.go")
	if should.NoError(err) {
		t.Log("upload test is pass")
	}
}
