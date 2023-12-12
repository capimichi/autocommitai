package Command

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use: "autocommitai",
	}

	cmd.AddCommand(AutoCommitCommand())
	cmd.AddCommand(CookieRefresh())

	return cmd
}
