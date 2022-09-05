package conf

import (
	"os"
)

var (
	OssAccessKeyId  string
	OssAccessSecret string
	OssEndPoint     string
	OssBucket       string
)

func init() {
	OssAccessKeyId = os.Getenv("OssAccessKeyId")
	OssAccessSecret = os.Getenv("OssAccessSecret")
	OssEndPoint = os.Getenv("OssEndPoint")
	OssBucket = os.Getenv("OssBucket")
}
