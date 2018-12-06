package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "edit",
	Long:  "edit",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("need 1 args.")
			return
		}

		editor := args[0]

		// if err := file.Backup(srcFile); err != nil {
		// 	msg := fmt.Sprintf("failed backup. err=%s", err)
		// 	log.Println(msg)
		// 	return
		// }

		fmt.Println(editor)
	},
}

func init() {
	RootCommand.AddCommand(editCommand)
	editCommand.Flags().StringP("owner", "o", "", "owner of remote file")
	editCommand.Flags().StringP("group", "g", "", "group of remote file")
	editCommand.Flags().StringP("mode", "m", "", "mode of remote file")
}
