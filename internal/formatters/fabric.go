package formatters

import "errors"

type FormatterFabric struct{}

func (f FormatterFabric) CreateFormatter(formatType string) (FormatterInterface, error) {
	switch formatType {
	case "stylish":
		return StylishFormatter{}, nil
	case "plain":
		return PlainFormatter{}, nil
	case "json":
		return JSONFormatter{}, nil
	default:
		return nil, errors.New("not implemented for format " + formatType)
	}
}
