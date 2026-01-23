package formatters

import "code/pkg/differ"

type FormatterInterface interface {
	Format(diff *differ.Diff, level int) (string, error)
}

type FormatterFabricInterface interface {
	CreateFormatter(formatType string) (FormatterInterface, error)
}
