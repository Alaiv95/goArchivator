package table

import (
	"strings"
)

type Generator interface {
	NewTable(text string) EncodingTable
}

type EncodingTable map[rune]string

type DecodingTree struct {
	Value rune
	Left  *DecodingTree
	Right *DecodingTree
}

func (et EncodingTable) Decode(text string) string {
	dt := et.DecodingTree()

	return dt.decode(text)
}

func (et EncodingTable) DecodingTree() DecodingTree {
	root := DecodingTree{}

	for k, bn := range et {
		root.add(bn, k)
	}

	return root
}

func (dt *DecodingTree) decode(code string) string {
	res := strings.Builder{}
	cur := dt

	for _, v := range code {
		switch v {
		case '1':
			cur = cur.Right
		case '0':
			cur = cur.Left
		}

		if cur == nil {
			break
		}

		if cur.Value != 0 {
			res.WriteRune(cur.Value)
			cur = dt
		}
	}

	return res.String()
}

func (dt *DecodingTree) add(binCode string, value rune) {
	cur := dt

	for _, v := range binCode {
		if v == '1' {
			if cur.Right == nil {
				cur.Right = &DecodingTree{}
			}
			cur = cur.Right
		}

		if v == '0' {
			if cur.Left == nil {
				cur.Left = &DecodingTree{}
			}
			cur = cur.Left
		}
	}

	if cur != nil && cur != dt {
		cur.Value = value
	}
}
