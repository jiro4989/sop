package command

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "edit",
	Long:  "edit",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Println("need 1 args.")
			return
		}

		// if err := file.Backup(srcFile); err != nil {
		// 	msg := fmt.Sprintf("failed backup. err=%s", err)
		// 	log.Println(msg)
		// 	return
		// }

		editor := args[0]
		srcFile := args[1]
		c := fmt.Sprintf("%s %s", editor, srcFile)
		if err := exec.Command(c).Run(); err != nil {
			log.Println("failed")
			return
		}

	},
}

func init() {
	RootCommand.AddCommand(editCommand)
	editCommand.Flags().StringP("owner", "o", "", "owner of remote file")
	editCommand.Flags().StringP("group", "g", "", "group of remote file")
	editCommand.Flags().StringP("mode", "m", "", "mode of remote file")
}
