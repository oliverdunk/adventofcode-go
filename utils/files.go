package utils

import (
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getRepositoryLocation() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath + "/src/github.com/oliverdunk/adventofcode-go/"
}

// ReadFile returns a string of the file's contents, assuming it exists
func ReadFile(path string) (string, error) {
	path = getRepositoryLocation() + path
	bytes, error := ioutil.ReadFile(path)

	if error != nil {
		return "", error
	}

	return string(bytes), nil
}

// SaveFile writes data to the specified file
func SaveFile(path string, data string) error {
	path = getRepositoryLocation() + path
	absolutePath, absolutePathError := filepath.Abs(path)

	if absolutePathError != nil {
		return absolutePathError
	}

	pathParts := strings.Split(absolutePath, "/")
	pathWithoutFile := strings.Join(pathParts[:len(pathParts)-1], "/")

	os.MkdirAll(pathWithoutFile, os.ModePerm)
	return ioutil.WriteFile(absolutePath, []byte(data), os.ModePerm)
}
