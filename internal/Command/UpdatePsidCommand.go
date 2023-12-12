package Command

import (
	"autocommitai/internal/Config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UpdatePsidCommand() *cobra.Command {

	psidKey := "psid"

	cmd := &cobra.Command{
		Use:   "update-psid",
		Short: "Update PSID",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := Config.NewDefaultConfig()
			err := defaultConfig.InitConfig()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			psid, err := cmd.Flags().GetString(psidKey)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if psid == "" {
				fmt.Println("Please enter your PSID:")
				_, err := fmt.Scanln(&psid)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			defaultConfig.SetPsid(psid)

			err = defaultConfig.Save()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().String(psidKey, "", "PSID")

	return cmd
}
