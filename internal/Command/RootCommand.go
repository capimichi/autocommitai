package Command

import (
	"autocommitai/internal/Config"
	"autocommitai/internal/Service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {

	defaultChoiceKey := "default-choice"
	ignoreUntrackedKey := "ignore-untracked"

	cmd := &cobra.Command{
		Use:   "autocommitai",
		Short: "AutoCommitAI is a tool to automatically commit your code",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := Config.NewDefaultConfig()
			err := defaultConfig.InitConfig()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			autoCommitAiService := Service.NewAutoCommitAiService()

			ignoreUntracked, err := cmd.Flags().GetBool(ignoreUntrackedKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			autoCommitAiService.SetIgnoreUntracked(ignoreUntracked)

			defaultChoice, err := cmd.Flags().GetString(defaultChoiceKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			autoCommitAiService.SetDefaultChoice(defaultChoice)

			err = autoCommitAiService.Execute()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().Bool(ignoreUntrackedKey, false, "Ignore untracked files when committing")
	cmd.Flags().String(defaultChoiceKey, "", "Default choice for commit message")

	return cmd
}
