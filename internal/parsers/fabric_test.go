package parsers_test

import (
	"code/internal/parsers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserFabric_CreateParser(t *testing.T) {
	tests := []struct {
		name       string
		formatType string
		want       parsers.ParserInterface
		wantErr    bool
	}{
		{"JSON parser", "json", parsers.JSONParser{}, false},
		{"YAML parser", "yml", parsers.YAMLParser{}, false},

		{"Unknown parser", "test", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f parsers.ParserFabric
			got, gotErr := f.CreateParser(tt.formatType)
			if tt.wantErr {
				assert.Error(t, gotErr, "not implemented for format "+tt.formatType)
			}
			assert.IsType(t, tt.want, got)
		})
	}
}
