package formatters

import (
	"code/pkg/differ"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/go-deepcopy"
)

func TestPlainFormatter_format(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		nodes   []*differ.DiffNode
		want    string
		wantErr bool
	}{
		{
			"plain diff",
			"plain",
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
					Type:          differ.OBJECT_TYPE,
				},
			},
			"../../testdata/fixture/expected/formatters_tests/plain.txt",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := differ.NewDiff()
			for _, node := range tt.nodes {
				d.AddNode(node)
			}
			var diffCopy differ.Diff
			_ = deepcopy.Copy(&diffCopy, d)
			d.AddNode(differ.NewNode("differ_field", &diffCopy))
			got, gotErr := PlainFormatter{}.format("", d)
			if tt.wantErr {
				assert.Errorf(t, gotErr, "")
			}
			expected, _ := os.ReadFile(tt.want)

			assert.Equal(t, string(expected), got)
		})
	}
}
