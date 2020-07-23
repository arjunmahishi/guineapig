package main

import (
	"encoding/json"
	"io/ioutil"
)

var config []route

func initConfig() error {
	if *response != "" {
		raw, err := ioutil.ReadFile(*response)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(raw, &config); err != nil {
			return err
		}
	}
	return nil
}
