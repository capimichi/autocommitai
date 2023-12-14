package Command

import (
	"autocommitai/internal/Config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UpdateBingCookieCommand() *cobra.Command {

	bingCookieKey := "bingCookie"

	cmd := &cobra.Command{
		Use:   "update-bingCookie",
		Short: "Update BingCookie",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := Config.NewDefaultConfig()
			err := defaultConfig.InitConfig()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			bingCookie, err := cmd.Flags().GetString(bingCookieKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if bingCookie == "" {
				fmt.Println("Please enter your BingCookie:")
				_, err := fmt.Scanln(&bingCookie)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			defaultConfig.SetBingCookie(bingCookie)

			err = defaultConfig.Save()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().String(bingCookieKey, "", "BingCookie")

	return cmd
}
