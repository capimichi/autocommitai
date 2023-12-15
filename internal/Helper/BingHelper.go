package Helper

import (
	"autocommitai/internal/Bing"
	"autocommitai/internal/Config"
)

type BingHelper struct {
	bing *Bing.DefaultBing
}

func NewBingHelper() *BingHelper {
	defaultConfig := Config.NewDefaultConfig()
	defaultBing := Bing.NewDefaultBing(defaultConfig.GetBingCookie())

	return &BingHelper{
		bing: defaultBing,
	}
}

func (bh *BingHelper) GetResponse(text string) (string, error) {
	//err := bh.bing.Ask(text)
	//if err != nil {
	//	return "", err
	//}

	response, err := bh.bing.GetResponse(text)
	if err != nil {
		return "", err
	}
	return response, nil
}
