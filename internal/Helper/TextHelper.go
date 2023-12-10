package Helper

import (
	"strings"
)

type TextHelper struct {}

func NewTextHelper() *TextHelper {
	return &TextHelper{}
}

func (th *TextHelper) ExtractJson(text string) string {
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")
	if start != -1 && end != -1 {
		return text[start : end+1]
	}
	return ""
}
