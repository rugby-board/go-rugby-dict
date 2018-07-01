package dict

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Dict object
type Dict struct {
	dictPath string
	wordList []Word
	size     int
}

// Word object
type Word struct {
	Origin      string `yaml:"w"`
	Translation string `yaml:"t"`
}

const (
	// RepoPath Relative path of this repo
	RepoPath = "github.com/rugby-board/go-rugby-dict"
	// DefaultDictPath of dict.yaml
	DefaultDictPath = "dict.yaml"
)

// LoadEnvConfPath ...
func loadRealPath(yamlPath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return "", errors.New("GOPATH is not set")
	}

	path := fmt.Sprintf("%s/src/%s/%s", gopath, RepoPath, yamlPath)
	return path, nil
}

// NewDict returns a Dict
func NewDict(dictPath string) *Dict {
	return &Dict{
		dictPath: dictPath,
		size:     0,
		wordList: make([]Word, 0, 1024),
	}
}

// NewDefaultDict returns a default Dict
func NewDefaultDict() *Dict {
	return &Dict{
		dictPath: DefaultDictPath,
		size:     0,
		wordList: make([]Word, 0, 1024),
	}
}

// Load a dict file
func (d *Dict) Load() error {
	realConfPath, err := loadRealPath(d.dictPath)
	if err != nil {
		log.Fatalf("Get real conf path failed: %#v", err)
		return err
	}
	yamlFile, err := ioutil.ReadFile(realConfPath)
	if err != nil {
		log.Fatalf("Read file failed: %#v", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &d.wordList)
	if err != nil {
		log.Fatalf("Unmarshal failed: %#v", err)
		return err
	}
	return nil
}

// Size returns the size of a dict
func (d *Dict) Size() int {
	d.size = len(d.wordList)
	return d.size
}

// Query English names
func (d *Dict) Query(englishName string) (string, error) {
	for _, word := range d.wordList {
		if word.Origin == englishName {
			return word.Translation, nil
		}
	}
	return "", errors.New("Word not found")
}
