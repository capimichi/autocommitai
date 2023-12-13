package Helper

import (
	"autocommitai/internal/Model"
	"encoding/json"
	"math/rand"
	"strings"
	"time"
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
		// fileName := strings.Split(file.GetPath(), "/")[0]
		return "Created file: " + file.GetPath(), nil
	}

	var commitMessage Model.CommitMessage
	var jsonPart string
	for i := 0; i < 3; i++ {
		prompt := ach.GetPrompt(file, diff)
		response, err := ach.BardHelper.GetResponse(prompt)

		if err == nil {
			jsonPart = ach.TextHelper.ExtractJson(response)

			if jsonPart != "" {
				err = json.Unmarshal([]byte(jsonPart), &commitMessage)

				if err == nil {
					break
				}
			}
		} else {
			// check if err message contains snlm0e (case insensitive)
			if strings.Contains(strings.ToLower(err.Error()), "snim0e") {
				return "", err
			}

		}
	}

	if err != nil {
		return "", err
	}

	return commitMessage.Message, nil
}

func (ach *AutoCommitHelper) GetPrompt(file Model.GitFile, diff string) string {
	possiblePrompts := []string{
		"Please generate a commit message for this, in this format { \"message\": \"commit message\" } \n",
		"Even if you are only an AI model, can you try to give a possible git commit message for this changes, in this format: { \"message\": \"your message here\" } \n",
		"Can you give a possible git commit message for this changes, in this format: { \"message\": \"your message here\" } \n",
		"Please, i need you to act as a human and give a possible git commit message for this changes, in this format: { \"message\": \"your message here\" } \n",
		"Can you try, even if you are an AI, to give a possible git commit message for this changes, in this format: { \"message\": \"your message here\" } \n",
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rand.Intn(len(possiblePrompts))
	prompt := possiblePrompts[index]
	prompt += "File: " + file.GetPath() + "\n"
	prompt += "---\n"
	prompt += diff + "\n"
	prompt += "---\n"
	return prompt
}
