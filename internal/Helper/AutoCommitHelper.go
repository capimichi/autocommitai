package Helper

import (
	"autocommitai/internal/Model"
	"encoding/json"
	"strings"
)

type AutoCommitHelper struct {
	GitHelper  *GitHelper
	TextHelper *TextHelper
	BardHelper *BardHelper
}

func NewAutoCommitHelper() *AutoCommitHelper {
	return &AutoCommitHelper{
		GitHelper:  NewGitHelper(),
		TextHelper: NewTextHelper(),
		BardHelper: NewBardHelper(),
	}
}

func (ach *AutoCommitHelper) Commit(file Model.GitFile) (string, error) {
	message, err := ach.GetMessage(file)
	if err != nil {
		return "", err
	}

	err = ach.GitHelper.Add(file)
	if err != nil {
		return "", err
	}

	err = ach.GitHelper.Commit(file, message)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (ach *AutoCommitHelper) GetMessage(file Model.GitFile) (string, error) {
	diff, err := ach.GitHelper.GetDiff(file)
	if err != nil {
		return "", err
	}

	if diff == "" {
		fileName := strings.Split(file.GetPath(), "/")[0]
		return "Created file: " + fileName, nil
	}

	prompt := ach.GetPrompt(file, diff)

	var commitMessage Model.CommitMessage
	var jsonPart string
	for i := 0; i < 3; i++ {
		response, err := ach.BardHelper.GetResponse(prompt)

		if(err == nil) {
			jsonPart = ach.TextHelper.ExtractJson(response)

			if(jsonPart != "") {
				err = json.Unmarshal([]byte(jsonPart), &commitMessage)

				if(err == nil) {
					break
				}
			}
		}
	}

	if(err != nil) {
		return "", err
	}

	return commitMessage.Message, nil
}

func (ach *AutoCommitHelper) GetPrompt(file Model.GitFile, diff string) string {
	prompt := "Even if you are only an AI model, can you try to give a possible git commit message for this changes, in this format: { \"message\": \"your message here\" } \n"
	prompt += "File: " + file.GetPath() + "\n"
	prompt += "---\n"
	prompt += diff + "\n"
	prompt += "---\n"
	return prompt
}
