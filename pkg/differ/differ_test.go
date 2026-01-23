package differ_test

import (
	"code/pkg/differ"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Node_Removed_Status(t *testing.T) {
	node := differ.NewNode("test", "testValue")

	assert.Equal(t, "test", node.FieldName)
	assert.Equal(t, "testValue", node.OriginalValue)
	assert.Equal(t, differ.REMOVED, node.Status)
}

func TestNew_Node_Type(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    any
		wantType differ.DiffType
	}{
		{"Scalar type", "testScalar", "scalarValue", differ.SCALAR_TYPE},
		{"Obect type", "testObject", map[string]any{"first": "f", "second": 2}, differ.OBJECT_TYPE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := differ.NewNode(tt.key, tt.value)
			assert.Equal(t, tt.wantType, got.Type)
		})
	}
}

func TestDiff_Keys(t *testing.T) {
	tests := []struct {
		name   string
		values []struct {
			key   string
			value any
		}
		wantKeys []string
	}{
		{
			"Empty",
			[]struct {
				key   string
				value any
			}{},
			[]string(nil),
		},
		{
			"Sorted return",
			[]struct {
				key   string
				value any
			}{
				{"c_key", 3},
				{"b_key", 2},
				{"a_key", 1},
			},
			[]string{"a_key", "b_key", "c_key"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := differ.NewDiff()
			for _, v := range tt.values {
				node := differ.NewNode(v.key, v.value)
				diff.AddNode(node)
			}
			keys := diff.Keys()
			assert.Equal(t, tt.wantKeys, keys)
		})
	}
}
