package parsers

import (
	"encoding/json"
)

type JSONParser struct{}

func (p JSONParser) Parse(content []byte) (map[string]any, error) {
	var res map[string]any

	err := json.Unmarshal(content, &res)
	
	if err != nil {
		return res, err
	}

	return res, nil
}