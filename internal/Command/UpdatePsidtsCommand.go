package Command

import (
	"autocommitai/internal/Config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UpdatePsidtsCommand() *cobra.Command {

	psidtsKey := "psidts"

	cmd := &cobra.Command{
		Use:   "update-psidts",
		Short: "Update PSIDTS",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := Config.NewDefaultConfig()
			err := defaultConfig.InitConfig()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			psidts, err := cmd.Flags().GetString(psidtsKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if psidts == "" {
				fmt.Println("Please enter your PSIDTS:")
				_, err := fmt.Scanln(&psidts)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			defaultConfig.SetPsidts(psidts)

			err = defaultConfig.Save()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().String(psidtsKey, "", "PSIDTS")

	return cmd
}
