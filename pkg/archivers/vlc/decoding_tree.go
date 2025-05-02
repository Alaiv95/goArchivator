package vlc

import "strings"

type encodingTable map[rune]string

type DecodingTree struct {
	Value rune
	Left  *DecodingTree
	Right *DecodingTree
}

func (et encodingTable) DecodingTree() DecodingTree {
	root := DecodingTree{}

	for k, bn := range et {
		root.add(bn, k)
	}

	return root
}

func (dt *DecodingTree) Decode(code string) string {
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

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
