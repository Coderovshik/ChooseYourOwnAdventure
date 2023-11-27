package main

import (
	"encoding/json"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func parseJSON(jsonBytes []byte) (arcs map[string]Arc, err error) {
	err = json.Unmarshal(jsonBytes, &arcs)
	return
}
