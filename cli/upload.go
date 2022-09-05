package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-oss/store"
	"go-oss/store/aliyun"
	"go-oss/store/minio"
	"go-oss/store/tencent"
)

var (
	ossProviderName string
	ossProvider     store.Storer
	ossBucketName   string
	uploadFileName  string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload file",
	Short:   "upload file",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch ossProviderName {
		case "aliyun":
			opts, _ := aliyun.NewDefaultAliOssStore()
			ossProvider, _ = aliyun.NewAliOssStore(opts)
		case "tencent":
			ossProvider, _ = tencent.NewTcTStore()
		case "minio":
			ossProvider, _ = minio.NewMiIOStore()
		default:
			return fmt.Errorf("not suppose oss storage provider")
		}
		ossProvider.Upload(ossBucketName, uploadFileName, uploadFileName)
		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProviderName, "provider", "p", "aliyun", "oss storage provider [tencent/minio]")
	f.StringVarP(&ossBucketName, "bucket_name", "b", "go-learn", "oss storage provider bucket name")
	f.StringVarP(&uploadFileName, "file_name", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
