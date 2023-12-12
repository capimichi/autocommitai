package Service

import (
	"autocommitai/internal/Config"
	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/all"
	"github.com/browserutils/kooky/browser/safari"
	"os"
)

type CookieRefreshService struct {
	browser string
}

func NewCookieRefreshService() *CookieRefreshService {
	return &CookieRefreshService{
		browser: "",
	}
}

func (c *CookieRefreshService) Execute() error {

	var cookies []*kooky.Cookie
	var err error
	if c.GetBrowser() == "safari" {
		dir, _ := os.UserHomeDir()
		cookiesFile := dir + "/Library/Cookies/Cookies.binarycookies"
		cookies, err = safari.ReadCookies(cookiesFile)
		if err != nil {
			return err
		}
	} else {
		cookies = kooky.ReadCookies(kooky.Valid)
	}

	psid := ""
	psidts := ""
	for _, cookie := range cookies {
		// check if name is "__Secure-1PSID"
		if cookie.Name == "__Secure-1PSID" {
			psid = cookie.Value
		}

		// check if name is "__Secure-1PSIDTS"
		if cookie.Name == "__Secure-1PSIDTS" {
			psidts = cookie.Value
		}
	}

	defaultConfig := Config.NewDefaultConfig()
	if psid != "" {
		defaultConfig.SetPsid(psid)
	}

	if psidts != "" {
		defaultConfig.SetPsidts(psidts)
	}

	err = defaultConfig.Save()
	if err != nil {
		return err
	}
	return nil
}

// GetBrowser returns the browser
func (c *CookieRefreshService) GetBrowser() string {
	return c.browser
}

// SetBrowser sets the browser
func (c *CookieRefreshService) SetBrowser(browser string) {
	c.browser = browser
}
