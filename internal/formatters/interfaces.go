package formatters

import "code/internal/differ"

type FormatterInterface interface {
	Format(diff *differ.Diff) (string, error)
}

type FormatterFabricInterface interface {
	CreateFormatter(formatType string) (FormatterInterface, error)
}
