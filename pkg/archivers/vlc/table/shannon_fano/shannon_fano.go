package shannon_fano

import (
	"archiver/pkg/archivers/vlc/table"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Generator struct{}

type charOccurrence map[rune]int
type encodingTable map[rune]code

type code struct {
	Value rune
	Count int
	Size  int
	Bit   int32
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	occurrences := countOccurrences(text)

	result := buildTable(occurrences)

	return result.export()
}

func (t encodingTable) export() table.EncodingTable {
	exportedTable := make(table.EncodingTable, len(t))

	for k, v := range t {
		binaryBits := fmt.Sprintf("%b", v.Bit)

		if diff := v.Size - len(binaryBits); diff > 0 {
			binaryBits = strings.Repeat("0", diff) + binaryBits
		}

		exportedTable[k] = binaryBits
	}

	return exportedTable
}

func buildTable(occurrence charOccurrence) encodingTable {
	codes := make([]code, 0, len(occurrence))

	for ch, cnt := range occurrence {
		codes = append(codes, code{
			Value: ch,
			Count: cnt,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Size != codes[j].Size {
			return codes[i].Size > codes[j].Size
		}

		return codes[i].Value < codes[j].Value
	})

	assignBits(codes)

	res := make(encodingTable, len(codes))

	for _, c := range codes {
		res[c.Value] = c
	}

	return res
}

func assignBits(codes []code) {
	if len(codes) < 2 {
		return
	}

	dividerIndex := bestDividerPos(codes)

	for i := range codes {
		// добавляем 1, левая часть
		// было 0000 0001
		// стало 0000 0010
		codes[i].Bit <<= 1
		codes[i].Size++

		if i >= dividerIndex {
			// добавляем 1, правая часть
			// было 0000 0010
			// стало 0000 0011
			codes[i].Bit |= 1
		}
	}

	assignBits(codes[:dividerIndex])
	assignBits(codes[dividerIndex:])
}

func bestDividerPos(codes []code) int {
	left := 0
	right := 0

	for _, c := range codes {
		right += c.Count
	}

	diff := right
	sliceIndex := 0

	for i, v := range codes {
		left += v.Count
		right -= v.Count
		tempDiff := int(math.Abs(float64(right) - float64(left)))

		if tempDiff >= diff {
			break
		}

		diff = tempDiff
		sliceIndex = i + 1
	}

	return sliceIndex
}

func countOccurrences(text string) charOccurrence {
	res := make(charOccurrence)

	for _, v := range text {
		res[v]++
	}

	return res
}
