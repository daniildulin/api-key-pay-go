package response

import "encoding/json"

type Requisite struct {
	Value       string      `json:"value"`
	Name        string      `json:"name"`
	Label       string      `json:"label"`
	Placeholder string      `json:"placeholder"`
	ID          string      `json:"id"`
	Tag         string      `json:"tag"`
	Type        string      `json:"type"`
	Hidden      json.Number `json:"hidden"`
	CheckEmpty  json.Number `json:"check_empty"`
	Psid        uint64      `json:"psid"`
	Readonly    uint8       `json:"readonly"`
}

type RequisiteShort struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
