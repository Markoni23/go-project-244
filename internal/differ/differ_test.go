package differ_test

import (
	"code/internal/differ"
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
			[]string{},
		},
		{
			"basic",
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
			assert.ElementsMatch(t, tt.wantKeys, keys)
		})
	}
}

func TestCalculateDiffTree(t *testing.T) {
	tests := []struct {
		name        string
		first       map[string]any
		second      map[string]any
		wantedNodes []*differ.DiffNode
	}{
		{
			"basic",
			map[string]any{"same": "same", "removed": "removed", "changed": "origChanged"},
			map[string]any{"same": "same", "added": "added", "changed": "newChanged", "object": map[string]any{"key": "val"}},
			[]*differ.DiffNode{
				{
					FieldName:     "added",
					NewValue:      "added",
					OriginalValue: nil,
					Status:        differ.ADDED,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName:     "changed",
					NewValue:      "newChanged",
					OriginalValue: "origChanged",
					Status:        differ.CHANGED,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName: "object",
					NewValue:  map[string]any{"key": "val"},
					Status:    differ.ADDED,
					Type:      differ.OBJECT_TYPE,
				},
				{
					FieldName:     "removed",
					OriginalValue: "removed",
					Status:        differ.REMOVED,
					Type:          differ.SCALAR_TYPE,
				},
				{
					FieldName:     "same",
					OriginalValue: "same",
					Status:        differ.SAME,
					Type:          differ.SCALAR_TYPE,
				},
			},
		},
		{
			"two array",
			map[string]any{"numbers": []int{1, 2, 3}},
			map[string]any{"numbers": []int{1, 2, 3, 4}},
			[]*differ.DiffNode{
				{
					FieldName:     "numbers",
					OriginalValue: []int{1, 2, 3},
					NewValue:      []int{1, 2, 3, 4},
					Status:        differ.CHANGED,
					Type:          differ.SCALAR_TYPE,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := differ.CalculateDiffTree(tt.first, tt.second)
			for _, key := range got.Keys() {
				assert.Contains(t, tt.wantedNodes, got.GetNode(key))
			}
		})
	}
}

func TestCalculateDiffTreeObjects(t *testing.T) {

	firstMap := map[string]any{"object": map[string]any{"numbers": []int{1, 2, 3}, "test": "value"}}
	secondMap := map[string]any{"object": map[string]any{"numbers": []int{1, 2, 3, 4}, "new_test": "new_value"}}

	nestedDiff := differ.NewDiff()
	nestedDiff.AddNode(&differ.DiffNode{
		FieldName:     "numbers",
		NewValue:      []int{1, 2, 3, 4},
		OriginalValue: []int{1, 2, 3},
		Status:        differ.CHANGED,
		Type:          differ.SCALAR_TYPE,
	})

	nestedDiff.AddNode(&differ.DiffNode{
		FieldName:     "test",
		OriginalValue: "value",
		Status:        differ.REMOVED,
		Type:          differ.SCALAR_TYPE,
	})
	nestedDiff.AddNode(&differ.DiffNode{
		FieldName: "new_test",
		NewValue:  "new_value",
		Status:    differ.ADDED,
		Type:      differ.SCALAR_TYPE,
	})

	diff := differ.NewDiff()

	diff.AddNode(&differ.DiffNode{
		FieldName:     "object",
		Status:        differ.SAME,
		Type:          differ.OBJECT_TYPE,
		OriginalValue: nestedDiff,
	})

	got, _ := differ.CalculateDiffTree(firstMap, secondMap)
	for _, key := range got.Keys() {
		node := got.GetNode(key)
		nestedDiff, isDiff := node.OriginalValue.(*differ.Diff)
		if isDiff {
			expectedNested, _ := diff.GetNode(key).OriginalValue.(*differ.Diff)
			for _, nestedKey := range nestedDiff.Keys() {
				assert.Equal(t, expectedNested.GetNode(nestedKey), nestedDiff.GetNode(nestedKey))
			}
		} else {
			assert.Equal(t, diff.GetNode(key), got.GetNode(key))
		}
	}
}
