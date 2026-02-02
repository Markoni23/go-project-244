package parsers_test

import (
	"code/internal/parsers"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected map[string]any
	}{
		{
			name:     "test one-level for json",
			filepath: "../../testdata/fixture/expected/parsers_tests/json.json",
			expected: map[string]any{"host": "hexlet.io", "timeout": float64(50), "proxy": "123.234.53.22", "follow": false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, _ := os.ReadFile(tt.filepath)
			res, err := parsers.JSONParser{}.Parse(content)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}
