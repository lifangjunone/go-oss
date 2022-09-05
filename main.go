package main

import (
	"fmt"
	"go-oss/conf"
	"go-oss/store/aliyun"
)

func main() {
	store := aliyun.Storer
	err := store.Upload(conf.OssBucket, "test.txt", "test.txt")
	if err != nil {
		fmt.Errorf("Upload file to oss is error: %s ", err)
	}
}
