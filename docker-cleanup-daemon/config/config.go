package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type CleanupSettings struct {
	StoppedContainers bool `yaml:"stopped_containers"`
	DanglingImages    bool `yaml:"dangling_images"`
	UnusedVolumes     bool `yaml:"unused_volumes"`
}

type ThresholdSettings struct {
	ContainersOlderThan string `yaml:"containers_older_than"`
}

type Config struct {
	Schedule   string            `yaml:"schedule"`
	Cleanup    CleanupSettings   `yaml:"cleanup"`
	Thresholds ThresholdSettings `yaml:"thresholds"`
	DryRun     bool              `yaml:"dry_run"`
}

// GetContainersOlderThanDuration parses the duration string or returns 0 if empty/invalid.
func (c *Config) GetContainersOlderThanDuration() time.Duration {
	if c.Thresholds.ContainersOlderThan == "" {
		return 0
	}
	d, err := time.ParseDuration(c.Thresholds.ContainersOlderThan)
	if err != nil {
		return 0
	}
	return d
}

// LoadConfig reads and decodes the YAML config file.
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
