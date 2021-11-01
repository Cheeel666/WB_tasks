package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config of service
type Config struct {
	Port         string `json:"port"`
	TestPort     string `json:"testPort"`
	MinID        int    `json:"minID"`
	BlockedUsers []int  `json:"BannedUsrs"`
}

// ParseConfig of service
func ParseConfig(configPath string) (*Config, error) {
	fileBody, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBody, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
