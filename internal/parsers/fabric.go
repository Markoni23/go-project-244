package parsers

import "errors"

type ParserFabric struct{}

func (f ParserFabric) CreateParser(extension string) (ParserInterface, error) {
	switch extension {
	case "json":
		return JSONParser{}, nil
	case "yaml", "yml":
		return YAMLParser{}, nil
	default:
		return nil, errors.New("not implemented for extension " + extension)
	}
}
