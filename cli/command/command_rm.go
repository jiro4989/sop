package command

import (
	"fmt"
	"log"
	"os"

	"github.com/jiro4989/sop/file"
	"github.com/spf13/cobra"
)

var rmCommand = &cobra.Command{
	Use:   "rm",
	Short: "rm",
	Long:  "rm",
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

		if err := os.Remove(srcFile); err != nil {
			msg := fmt.Sprintf("failed remove. err=%s", err)
			log.Println(msg)
			return
		}
	},
}

func init() {
	RootCommand.AddCommand(rmCommand)
	rmCommand.Flags().StringP("owner", "o", "", "owner of remote file")
	rmCommand.Flags().StringP("group", "g", "", "group of remote file")
	rmCommand.Flags().StringP("mode", "m", "", "mode of remote file")
}
