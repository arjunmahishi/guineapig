package main

import "encoding/json"

type route struct {
	Method             string          `json:"method"`
	Path               string          `json:"path"`
	ResponseStatusCode int             `json:"response_status_code"`
	ResponseBody       json.RawMessage `json:"response_body"`
}
