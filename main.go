package main

import (
	"fmt"
	"go-oss/store/aliyun"
)

func main() {
	//if err := cli.RootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//}
	opts, _ := aliyun.NewDefaultAliOssStore()
	ossProvider, _ := aliyun.NewAliOssStore(opts)
	err := ossProvider.Upload("go-learn", "main.go", "main.go")
	fmt.Println(err)
}
