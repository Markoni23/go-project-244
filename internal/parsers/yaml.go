package parsers

import "gopkg.in/yaml.v3"

type YAMLParser struct{}

func (y YAMLParser) Parse(content []byte) (map[string]any, error) {
	var res map[string]any

	err := yaml.Unmarshal(content, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
