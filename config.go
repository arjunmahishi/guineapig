package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Config map[string]map[string]struct {
	StatusCode int             `json:"status_code"`
	Response   json.RawMessage `json:"response"`
}

func InitConfig(configFile string) error {
	if configFile != "" {
		raw, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(raw, &Config); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("config file not provided")
}
