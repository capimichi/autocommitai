package Command

import (
	"autocommitai/internal/Config"
	"autocommitai/internal/Service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CookieRefresh() *cobra.Command {

	defaultBrowserKey := "default-browser"

	cmd := &cobra.Command{
		Use:   "cookie-refresh",
		Short: "Refresh cookie",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := Config.NewDefaultConfig()
			err := defaultConfig.InitConfig()
			cookieRefreshService := Service.NewCookieRefreshService()

			defaultBrowser, err := cmd.Flags().GetString(defaultBrowserKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			cookieRefreshService.SetBrowser(defaultBrowser)

			err = cookieRefreshService.Execute()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().String(defaultBrowserKey, "", "Default browser (chrome, firefox, safari, opera)")

	return cmd
}
