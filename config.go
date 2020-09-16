package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var config map[string]map[string]struct {
	StatusCode int             `json:"status_code"`
	Response   json.RawMessage `json:"response"`
}

func initConfig() error {
	if *configFile != "" {
		raw, err := ioutil.ReadFile(*configFile)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(raw, &config); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("config file not provided")
}
