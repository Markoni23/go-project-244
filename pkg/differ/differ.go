package differ

import (
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
		nodes: make(map[string]*DiffNode, 0),
		keys:  make([]string, 0),
	}
}

func (d *Diff) Keys() []string {
	return slices.Sorted(slices.Values(d.keys))
}

func (d *Diff) AddNode(node *DiffNode) {
	d.nodes[node.FieldName] = node
	d.keys = append(d.keys, node.FieldName)
}

func (d *Diff) CheckNode(key string, value any) {
	node, exists := d.nodes[key]

	if !exists {
		newNode := DiffNode{Type: ValueType(value), Status: ADDED, NewValue: value, FieldName: key}
		d.AddNode(&newNode)
		return
	}

	if node.Type == OBJECT_TYPE && ValueType(value) == OBJECT_TYPE {
		node.Status = SAME
		node.OriginalValue = CalculateDiffTree(node.OriginalValue.(map[string]any), value.(map[string]any))
		return
	}

	if node.OriginalValue == value {
		node.Status = SAME
	} else {
		node.Status = CHANGED
		node.NewValue = value
	}
}

func (d *Diff) GetNode(key string) *DiffNode {
	return d.nodes[key]
}

func CalculateDiffTree(first, second map[string]any) *Diff {
	diff := NewDiff()

	for key, value := range first {
		node := NewNode(key, value)
		diff.AddNode(node)
	}

	for key, value := range second {
		diff.CheckNode(key, value)
	}

	return diff
}
