package command

import (
	"fmt"
	"log"

	"github.com/jiro4989/sop/file"
	"github.com/spf13/cobra"
)

var cpCommand = &cobra.Command{
	Use:   "cp",
	Short: "cp copies file to remote server",
	Long:  "cp copies file to remote server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Println("need 2 args.")
			return
		}

		srcFile := args[0]
		dstFile := args[1]

		if err := file.Backup(srcFile); err != nil {
			msg := fmt.Sprintf("failed backup. err=%s", err)
			log.Println(msg)
			return
		}

		if err := file.Copy(srcFile, dstFile); err != nil {
			msg := fmt.Sprintf("failed copy. err=%s", err)
			log.Println(msg)
			return
		}
	},
}

func init() {
	RootCommand.AddCommand(cpCommand)
	cpCommand.Flags().StringP("owner", "o", "", "owner of remote file")
	cpCommand.Flags().StringP("group", "g", "", "group of remote file")
	cpCommand.Flags().StringP("mode", "m", "", "mode of remote file")
}