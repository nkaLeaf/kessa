package handlers

import (
	"encoding/json"

	"os"
)

type Conf struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

var Cfg Conf

func LoadConf() (Conf, error) {

	data, err := os.ReadFile("conf.json")
	Catch("config read", "fatal", err)
	err = json.Unmarshal(data, &Cfg)
	return Cfg, err
}
