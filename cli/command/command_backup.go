package command

import (
	"fmt"
	"log"

	"github.com/jiro4989/sop/file"
	"github.com/spf13/cobra"
)

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "backup copies file to remote server",
	Long:  "backup copies file to remote server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("need 1 args.")
			return
		}

		srcFile := args[0]

		if err := file.Backup(srcFile); err != nil {
			msg := fmt.Sprintf("failed backup. err=%s", err)
			log.Println(msg)
			return
		}
	},
}

func init() {
	RootCommand.AddCommand(backupCommand)
}
