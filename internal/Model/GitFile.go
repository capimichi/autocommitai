package Model

import (
	"os"
	"strings"
)

type GitFile struct {
	Path   string
	Status string
}

func NewGitFile(path string, status string) *GitFile {
	return &GitFile{Path: path, Status: status}
}

func (gf *GitFile) IsDir() bool {
	filePath := gf.Path
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func (gf *GitFile) IsTracked() bool {
	return gf.Status != "?" && gf.Status != "??"
}

func (gf *GitFile) GetPath() string {
	return gf.Path
}

func (gf *GitFile) GetStatus() string {
	v := strings.TrimSpace(gf.Status)
	return v
}
