package response

import "encoding/json"

type DirectionSuccessResponse struct {
	Status string      `json:"status"`
	Value  []Direction `json:"value"`
}

type Direction struct {
	OutValute string      `json:"out_valute"`
	Reserve   json.Number `json:"reserve"`
	OutGroup  string      `json:"out_group"`
	InGroup   string      `json:"in_group"`
	OutName   string      `json:"out_name"`
	City      string      `json:"city"`
	In        json.Number `json:"in"`
	InValute  string      `json:"in_valute"`
	InName    string      `json:"in_name"`
	Out       json.Number `json:"out"`
	Psid2     int         `json:"psid2"`
	ID        uint64      `json:"id"`
	InMin     float64     `json:"in_min"`
	InMax     float64     `json:"in_max"`
	OutMin    float64     `json:"out_min"`
	Psid1     int         `json:"psid1"`
	Direct    int         `json:"direct"`
	Enabled   uint8       `json:"enabled"`
}
