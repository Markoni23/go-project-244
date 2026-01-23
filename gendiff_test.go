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
		format     string
		want       string
		wantErr    bool
	}{

		{
			"json -> stylish",
			"testdata/fixture/json/nested1.json",
			"testdata/fixture/json/nested2.json",
			"stylish",
			"testdata/fixture/expected/nested_stylish.txt",
			false,
		},
		{
			"yml -> stylish",
			"testdata/fixture/yml/nested1.yml",
			"testdata/fixture/yml/nested2.yml",
			"stylish",
			"testdata/fixture/expected/nested_stylish.txt",
			false,
		},
		{
			"json -> plain",
			"testdata/fixture/json/nested1.json",
			"testdata/fixture/json/nested2.json",
			"plain",
			"testdata/fixture/expected/nested_plain.txt",
			false,
		},
		{
			"yml -> plain",
			"testdata/fixture/yml/nested1.yml",
			"testdata/fixture/yml/nested2.yml",
			"plain",
			"testdata/fixture/expected/nested_plain.txt",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected, _ := os.ReadFile(tt.want)

			got, gotErr := code.GenDiff(tt.firstFile, tt.secondFile, tt.format)
			if tt.wantErr {
				assert.NotNil(t, gotErr, "expected error")
			}

			assert.Equal(t, string(expected), got)
		})
	}
}
