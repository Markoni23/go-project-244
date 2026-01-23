package formatters

import (
	"code/pkg/differ"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylishFormatter_formatDiff(t *testing.T) {
	tests := []struct {
		name    string
		nodes   []*differ.DiffNode
		want    string
		wantErr bool
	}{
		{"Empty diff", []*differ.DiffNode{}, "../../testdata/fixture/expected/formatters_tests/empty_stylish.txt", false},
		{
			"basic diff",
			[]*differ.DiffNode{
				{
					FieldName:     "same_field",
					OriginalValue: "same",
					Status:        differ.SAME,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName: "added_field",
					NewValue:  "added",
					Status:    differ.ADDED,
					Type:      differ.SCALAR_TYPE,
				},
				{
					FieldName:     "removed_field",
					OriginalValue: "removed",
					Status:        differ.REMOVED,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName:     "changed_field",
					OriginalValue: "orig",
					NewValue:      "new",
					Status:        differ.CHANGED,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName:     "object_field",
					OriginalValue: map[string]any{"test": "t"},
					Status:        differ.SAME,
					Type:          differ.SCALAR_TYPE,
				},
			},
			"../../testdata/fixture/expected/formatters_tests/stylish.txt",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := differ.NewDiff()
			for _, node := range tt.nodes {
				d.AddNode(node)
			}

			got, gotErr := StylishFormatter{}.formatDiff(d, 0)
			if tt.wantErr {
				assert.Errorf(t, gotErr, "")
			}
			expected, _ := os.ReadFile(tt.want)

			assert.Equal(t, string(expected), got)
		})
	}
}
