package command

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "edit",
	Long:  "edit",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start edit. args=", args)
		if len(args) < 2 {
			log.Println("need 1 args.", args)
			return
		}

		// if err := file.Backup(srcFile); err != nil {
		// 	msg := fmt.Sprintf("failed backup. err=%s", err)
		// 	log.Info(msg)
		// 	return
		// }

		editor := args[0]
		srcFile := args[1]
		log.Println("editor=", editor, ", file=", srcFile)

		c := exec.Command(editor, srcFile)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin

		if err := c.Run(); err != nil {
			log.Println("failed exec command. editor=", editor, ", srcFile-", srcFile, ", err=", err)
			return
		}

		log.Println("end edit.")
	},
}

func init() {
	RootCommand.AddCommand(editCommand)
	editCommand.Flags().StringP("owner", "o", "", "owner of remote file")
	editCommand.Flags().StringP("group", "g", "", "group of remote file")
	editCommand.Flags().StringP("mode", "m", "", "mode of remote file")
}
