package Helper

import (
	"autocommitai/internal/Model"
	"os/exec"
	"strings"
)

type GitHelper struct{}

func NewGitHelper() *GitHelper {
	return &GitHelper{}
}

func (gh *GitHelper) GetAddedFiles() ([]string, error) {
	out, err := exec.Command("git", "diff", "--name-only", "--cached").Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}

func (gh *GitHelper) GetModifiedFiles() ([]string, error) {
	out, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}

func (gh *GitHelper) GetStatusFiles() ([]Model.GitFile, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return nil, err
	}
	paths := strings.Split(string(out), "\n")
	// remove empty lines
	var filteredPaths []string
	for _, path := range paths {
		if len(path) > 0 {
			filteredPaths = append(filteredPaths, path)
		}
	}
	paths = filteredPaths
	var filePaths []Model.GitFile
	for _, path := range paths {
		filePath := strings.TrimSpace(path[3:])
		if strings.Contains(filePath, "->") {
			parts := strings.Split(filePath, "->")
			filePath = strings.TrimSpace(parts[len(parts)-1])
		}

		status := strings.TrimSpace(path[0:2])
		gitFile := Model.GitFile{Path: filePath, Status: status}
		filePaths = append(filePaths, gitFile)
	}
	return filePaths, nil
}

func (gh *GitHelper) GetDiff(gitFile Model.GitFile) (string, error) {
	out, err := exec.Command("git", "diff", gitFile.GetPath()).Output()
	if err != nil {
		return "", err
	}
	if len(out) == 0 {
		out, err = exec.Command("git", "diff", "--cached", gitFile.GetPath()).Output()
		if err != nil {
			return "", err
		}
	}
	return string(out), nil
}

func (gh *GitHelper) Add(gitFile Model.GitFile) error {
	return exec.Command("git", "add", gitFile.GetPath()).Run()
}

func (gh *GitHelper) Commit(gitFile Model.GitFile, message string) error {
	return exec.Command("git", "commit", "-m", message, gitFile.GetPath()).Run()
}
