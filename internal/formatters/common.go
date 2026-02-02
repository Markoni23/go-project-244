package formatters

import (
	"code/internal/differ"
)

func isDiff(value any) bool {
	if value == nil {
		return false
	}

	_, ok := value.(*differ.Diff)
	return ok
}
