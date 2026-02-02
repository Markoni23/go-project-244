package differ

import (
	"fmt"
	"reflect"
	"slices"
)

type DiffStatus uint8

type DiffType uint8

const (
	SAME DiffStatus = iota
	CHANGED
	REMOVED
	ADDED
)

const (
	OBJECT_TYPE DiffType = iota
	SCALAR_TYPE
)

type DiffNode struct {
	FieldName               string
	OriginalValue, NewValue any
	Status                  DiffStatus
	Type                    DiffType
}

func NewNode(key string, value any) *DiffNode {
	res := DiffNode{FieldName: key, Status: REMOVED, Type: ValueType(value), OriginalValue: value}
	return &res
}

func ValueType(value any) DiffType {
	if value == nil {
		return SCALAR_TYPE
	}
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Map {
		return OBJECT_TYPE
	} else {
		return SCALAR_TYPE
	}
}

type Diff struct {
	nodes map[string]*DiffNode
	keys  []string
}

func NewDiff() *Diff {
	return &Diff{
		nodes: make(map[string]*DiffNode),
		keys:  make([]string, 0),
	}
}

func (d *Diff) Keys() []string {
	return d.keys
}

func (d *Diff) AddNode(node *DiffNode) {
	d.nodes[node.FieldName] = node
	d.keys = append(d.keys, node.FieldName)
}

func (d *Diff) CheckNode(key string, value any) error {
	node, exists := d.nodes[key]

	if !exists {
		newNode := DiffNode{Type: ValueType(value), Status: ADDED, NewValue: value, FieldName: key}
		d.AddNode(&newNode)
		return nil
	}

	if node.Type == OBJECT_TYPE && ValueType(value) == OBJECT_TYPE {
		node.Status = SAME
		origVal, ok := node.OriginalValue.(map[string]any)
		if !ok {
			return fmt.Errorf("value %v couldn't be converted to map[string]any type", node.OriginalValue)
		}

		newVal, ok := value.(map[string]any)
		if !ok {
			return fmt.Errorf("value %v couldn't be converted to map[string]any type", value)
		}

		calculatedDiff, err := CalculateDiffTree(origVal, newVal)
		if err != nil {
			return err
		}
		node.OriginalValue = calculatedDiff
		return nil
	}

	if reflect.DeepEqual(node.OriginalValue, value) {
		node.Status = SAME
	} else {
		node.Status = CHANGED
		node.NewValue = value
	}
	return nil
}

func (d *Diff) GetNode(key string) *DiffNode {
	return d.nodes[key]
}

func CalculateDiffTree(first, second map[string]any) (*Diff, error) {
	diff := NewDiff()

	for key, value := range first {
		node := NewNode(key, value)
		diff.AddNode(node)
	}

	for key, value := range second {
		if err := diff.CheckNode(key, value); err != nil {
			return diff, err
		}
	}

	return diff, nil
}

func SortedKeys(diff *Diff) []string {
	return slices.Sorted(slices.Values(diff.keys))
}
