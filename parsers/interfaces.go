package parsers

type ParserInterface interface {
	Parse(content []byte) (map[string]any, error)
}
