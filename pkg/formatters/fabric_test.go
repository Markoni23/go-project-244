package formatters_test

import (
	"code/pkg/formatters"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatterFabric_CreateFormatter(t *testing.T) {
	tests := []struct {
		name       string
		formatType string
		want       formatters.FormatterInterface
		wantErr    bool
	}{
		{"Stylish formatter", "stylish", formatters.StylishFormatter{}, false},
		{"Unknown formatter", "test", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f formatters.FormatterFabric
			got, gotErr := f.CreateFormatter(tt.formatType)
			if tt.wantErr {
				assert.Error(t, gotErr, "not implemented for format "+tt.formatType)
			}
			assert.IsType(t, tt.want, got)
		})
	}
}
