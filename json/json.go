package json

import (
	"encoding/json"
)

// func basics() {

// }

type Metric struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Value interface{} `json:"value"` // can hold any data type (bool, float, string) [numbers parsed as float64]
}

func Marshal(m Metric) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(b), err
}

func Unmarshal(data string) ([]Metric, error) {
	var res []Metric

	if err := json.Unmarshal([]byte(data), &res); err != nil {
		return res, err
	}

	return res, nil
}
