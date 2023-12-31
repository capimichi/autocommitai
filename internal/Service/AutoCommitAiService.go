package Service

import (
	"autocommitai/internal/Helper"
	"autocommitai/internal/Model"
	"fmt"
)

type AutoCommitAiService struct {
	ignoreUntracked bool
	defaultChoice   string
}

func NewAutoCommitAiService() *AutoCommitAiService {
	return &AutoCommitAiService{
		ignoreUntracked: false,
		defaultChoice:   "",
	}
}

func (ac *AutoCommitAiService) Execute() error {

	gitHelper := Helper.NewGitHelper()
	autoCommitHelper := Helper.NewAutoCommitHelper()

	files, err := gitHelper.GetStatusFiles()
	if err != nil {
		return err
	}

	if ac.GetIgnoreUntracked() {
		filterFiles := make([]Model.GitFile, 0)
		for _, file := range files {
			if file.IsTracked() {
				filterFiles = append(filterFiles, file)
			}
		}
		files = filterFiles
	}

	for _, file := range files {

		// ask the user what to do (1. commit, 2. ignore, 3. skip)
		fmt.Println("What do you want to do with this file: " + file.GetPath() + "?")

		var choice string
		if len(ac.GetDefaultChoice()) <= 0 {
			if !file.IsDir() {
				fmt.Println("1. Commit with auto message")
			}
			fmt.Println("2. Commit with custom message")
			fmt.Println("3. Ignore")
			fmt.Println("4. Skip")

			_, err = fmt.Scanln(&choice)
			if err != nil {
				return err
			}
		} else {
			choice = ac.defaultChoice
		}

		// switch on the choice
		switch choice {
		case "1":
			if !file.IsDir() {
				fmt.Println("Committing with auto message...")
				var message string
				message, err = autoCommitHelper.Commit(file)
				if err != nil {

					return err
				}
				fmt.Println("")
				fmt.Println("Committed with message: " + message)
				fmt.Println("")
			}
		}

		if choice == "2" {
			fmt.Println("Committing with custom message...")
			var message string
			// ask message
			fmt.Println("Please enter your commit message:")
			_, err = fmt.Scanln(&message)
			if err != nil {
				return err
			}

			err = gitHelper.Commit(file, message)
			if err != nil {
				return err
			}
			fmt.Println("")
		}

	}

	if len(files) <= 0 {
		fmt.Println("No files to commit")
	}

	return nil

}

func (a *AutoCommitAiService) GetIgnoreUntracked() bool {
	return a.ignoreUntracked
}

func (a *AutoCommitAiService) SetIgnoreUntracked(ignoreUntracked bool) {
	a.ignoreUntracked = ignoreUntracked
}

func (a *AutoCommitAiService) GetDefaultChoice() string {
	return a.defaultChoice
}

func (a *AutoCommitAiService) SetDefaultChoice(defaultChoice string) {
	a.defaultChoice = defaultChoice
}
