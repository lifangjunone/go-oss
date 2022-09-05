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
		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProviderName, "provider", "p", "aliyun", "oss storage provider [tencent/minio]")
	RootCmd.AddCommand(UploadCmd)
}
