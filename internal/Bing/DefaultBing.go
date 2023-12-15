package Bing

import (
	"autocommitai/internal/Python/Bing/data"
	"bytes"
	"github.com/kluctl/go-embed-python/embed_util"
	"strings"

	"github.com/kluctl/go-embed-python/python"
)

type DefaultBing struct {
	bingCookie string
}

func NewDefaultBing(bingCookie string) *DefaultBing {
	return &DefaultBing{
		bingCookie: bingCookie,
	}
}

func (db *DefaultBing) GetResponse(text string) (string, error) {
	name := "defaultBing"
	ep, err := python.NewEmbeddedPython(name)
	if err != nil {
		return "", err
	}

	var embeddedFiles *embed_util.EmbeddedFiles
	embeddedFiles, err = embed_util.NewEmbeddedFiles(data.Data, name)
	if err != nil {
		return "", err
	}
	ep.AddPythonPath(embeddedFiles.GetExtractedPath())

	prompt := text
	// replace new line with space
	prompt = strings.Replace(prompt, "\\", " ", -1)
	// replace new line with space
	prompt = strings.Replace(prompt, "\n", " ", -1)
	// replace ' with empty \'
	prompt = strings.Replace(prompt, "'", "", -1)

	//set prompt max length to 500
	if len(prompt) > 500 {
		prompt = prompt[:500]
	}

	cmd := ep.PythonCmd("-c", `
import asyncio
from sydney import SydneyClient

async def main():
    bing_cookies = "`+db.bingCookie+`"
    sydney = SydneyClient()
    await sydney.start_conversation()
    response = await sydney.ask('`+prompt+`', search=False)
    print(response)
    await sydney.close_conversation()

if __name__ == '__main__':
    asyncio.run(main())
`)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	// After running the command, you can get the output as a string
	outStr, errStr := stdout.String(), stderr.String()

	if err != nil {
		return errStr, err
	}

	return outStr, nil
}
