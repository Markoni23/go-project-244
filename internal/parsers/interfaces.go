package parsers

type ParserInterface interface {
	Parse(content []byte) (map[string]any, error)
}

type ParserFabricInterface interface {
	CreateParser(extension string) (ParserInterface, error)
}
