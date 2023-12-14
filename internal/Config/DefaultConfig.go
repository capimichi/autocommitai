package Config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DefaultConfig struct {
}

func NewDefaultConfig() *DefaultConfig {
	return &DefaultConfig{}
}

func (c *DefaultConfig) InitConfig() error {

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

		fmt.Println("Please enter your BingCookie:")
		var bingCookie string
		_, err := fmt.Scanln(&bingCookie)
		if err != nil {
			return err
		}
		viper.Set(c.GetBingCookieConfigKey(), bingCookie)

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

func (c *DefaultConfig) GetAppName() string {
	return "autocommitai"
}

func (c *DefaultConfig) GetBingCookieConfigKey() string {
	return "bingCookie"
}

func (c *DefaultConfig) GetBingCookie() string {
	return viper.GetString(c.GetBingCookieConfigKey())
}

func (c *DefaultConfig) SetBingCookie(bingCookie string) {
	viper.Set(c.GetBingCookieConfigKey(), bingCookie)
}

func (c *DefaultConfig) Save() error {
	return viper.WriteConfig()
}
