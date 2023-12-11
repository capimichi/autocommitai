package Helper

import (
	"autocommitai/internal/Config"
	bard "github.com/aquasecurity/gobard"
)

type BardHelper struct {
	bard *bard.Bard
}

func NewBardHelper() *BardHelper {
	defaultConfig := Config.NewDefaultConfig()
	defaultBard := bard.New(defaultConfig.GetPsid(), defaultConfig.GetPsidts())

	return &BardHelper{
		bard: defaultBard,
	}
}

func (bh *BardHelper) GetResponse(text string) (string, error) {
	err := bh.bard.Ask(text)
	if err != nil {
		return "", err
	}

	response := bh.bard.GetAnswer()
	return response, nil
}
