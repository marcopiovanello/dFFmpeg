package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type config struct {
	Root          string `yaml:"root"`
	FFMpegPath    string `yaml:"ffmpegPath"`
	ReleasePrefix string `yaml:"releasePrefix"`
}

var (
	instance     *config
	instanceOnce sync.Once
)

func LoadFile(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}

	instance = &config{}

	return yaml.NewDecoder(fd).Decode(instance)
}

func Instance() *config {
	if instance == nil {
		instanceOnce.Do(func() {
			instance = &config{}
		})
	}

	return instance
}
