package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Server struct {
	IPAddress string `json:"IPAddress"`
	Port      int    `json:"Port"`
}

type Configs struct {
	Accounts string `json:"Accounts"`
}

type Game struct {
	ShardName string  `json:"ShardName"`
	Configs   Configs `json:"Configs"`
}

type Config struct {
	Server Server `json:"Server"`
	Game   Game   `json:"Game"`
}

func LoadConfiguration(path string) Config {
	wd, _ := os.Getwd()
	wd = filepath.Dir(wd)
	fpath := wd + path
	fmt.Println("Data: Loading configuration from: " + fpath + "...")
	data, err := os.Open(fpath)
	if err != nil {
		panic("Data: Unable to Load configuration data: " + err.Error())
	}
	defer data.Close()

	var config Config
	decoder := json.NewDecoder(data)
	err = decoder.Decode(&config)
	if err != nil {
		panic("Data: Unable to Decode configuration data: " + err.Error())
	}

	return config
}

func (c *Config) GetConfigPath(name string) (string, error) {
	switch name {
	case "Accounts":
		wd, _ := os.Getwd()
		wd = filepath.Dir(wd)
		return wd + c.Game.Configs.Accounts, nil
	}
	return "", fmt.Errorf("Data: Unable to find config path for: " + name)
}
