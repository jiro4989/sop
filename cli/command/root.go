package command

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "sop",
	Short: "sop is safety operation",
	Long:  "sop is safety operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()
}
