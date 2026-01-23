package code_test

import (
	"code"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name       string
		firstFile  string
		secondFile string
		want       string
		wantErr    bool
	}{

		{
			"plain json",
			"testdata/fixture/json/nested1.json",
			"testdata/fixture/json/nested2.json",
			"testdata/fixture/expected/nested_stylish.txt",
			false,
		},
		{
			"Plain yml -> json format",
			"testdata/fixture/yml/nested1.yml",
			"testdata/fixture/yml/nested2.yml",
			"testdata/fixture/expected/nested_stylish.txt",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected, _ := os.ReadFile(tt.want)

			got, gotErr := code.GenDiff(tt.firstFile, tt.secondFile, "stylish")
			if tt.wantErr {
				assert.NotNil(t, gotErr, "expected error")
			}

			assert.Equal(t, string(expected), got)
		})
	}
}
