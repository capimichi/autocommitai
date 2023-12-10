package Config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type DefaultConfig struct {

}

func NewDefaultConfig () *DefaultConfig {
	return &DefaultConfig{}
}

func (c *DefaultConfig) InitConfig () error {

	appName := c.GetAppName()

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := userHomeDir + "/." + appName
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {

		fmt.Println("No configuration file found.")

		fmt.Println("Please enter your PSID:")
		var psid string
		_, err := fmt.Scanln(&psid)
		if err != nil {
			return err
		}
		viper.Set(c.GetPsidConfigKey(), psid)

		fmt.Println("Please enter your PSIDTS:")
		var psidts string
		_, err = fmt.Scanln(&psidts)
		if err != nil {
			return err
		}
		viper.Set(c.GetPsidtsConfigKey(), psidts)

		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err = os.Mkdir(configDir, 0755)
			if err != nil {
				return err
			}
		}

		// create config file
		_, err = os.Create(configDir + "/config.json")

		err = viper.WriteConfig()
		if err != nil {
			return err
		}

	}

	return nil
}

func (c *DefaultConfig) GetAppName () string {
	return "autocommitai"
}

func (c *DefaultConfig) GetPsidConfigKey () string {
	return "psid"
}

func (c *DefaultConfig) GetPsidtsConfigKey () string {
	return "psidts"
}

func (c *DefaultConfig) GetPsid () string {
	return viper.GetString(c.GetPsidConfigKey())
}

func (c *DefaultConfig) GetPsidts () string {
	return viper.GetString(c.GetPsidtsConfigKey())
}