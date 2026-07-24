package handlers

import (
	"encoding/json"

	"os"
)

type Conf struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func LoadConf() (Conf, error) {
	var cfg Conf
	data, err := os.ReadFile("conf.json")
	Catch("config read", "fatal", err)
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
