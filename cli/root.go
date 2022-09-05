package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version bool
)

var RootCmd = &cobra.Command{
	Use:     "oss-go-cli",
	Long:    "oss-go-cli usage example",
	Short:   "oss go example",
	Example: "oss-go-cli cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("Version: v1.0.0")
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "V", false, "oss-go version:")
}
