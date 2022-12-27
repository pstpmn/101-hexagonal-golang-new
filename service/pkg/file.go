package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type IFile interface {
	Read(path string) (string, error)
	IsFileExists(path string) bool
	CreateFile(content string, path string) error
}

type file struct {
}

func (f file) Read(path string) (string, error) {
	absPath, _ := filepath.Abs(path)
	data, err := os.ReadFile(absPath)
	return string(data), err
}

func (f file) IsFileExists(path string) bool {
	absPath, _ := filepath.Abs(path)
	if _, err := os.Stat(absPath); err == nil {
		return true
	} else {
		return false
	}
}

func (f file) CreateFile(content string, path string) error {
	absPath, _ := filepath.Abs(path)
	data := []byte(content)
	err := ioutil.WriteFile(absPath, data, 777)

	if err != nil {
		return err
	}
	return nil
}

func NewFile() IFile {
	return &file{}
}
