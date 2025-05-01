package chunks

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk
type BinaryChunk string

type HexChunks []HexChunk

type HexChunk string

func ToBinaryChunks(text string, chunkSize int) BinaryChunks {
	strlen := utf8.RuneCountInString(text)
	chunksCnt := strlen / chunkSize
	res := make(BinaryChunks, 0, chunksCnt)

	buf := strings.Builder{}

	if strlen%chunkSize != 0 {
		chunksCnt++
	}

	for i, v := range text {
		buf.WriteRune(v)

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

func (h HexChunks) ToString() string {
	const sep = " "

	switch len(h) {
	case 0:
		return ""
	case 1:
		return string(h[0])
	}

	buf := strings.Builder{}
	buf.WriteString(string(h[0]))

	for _, v := range h[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(v))
	}
	return buf.String()
}

func (cs BinaryChunks) ToHex(chunkSize int) HexChunks {
	res := make(HexChunks, 0, len(cs))

	for _, c := range cs {
		res = append(res, c.toHex(chunkSize))
	}

	return res
}

func (c BinaryChunk) toHex(chunkSize int) HexChunk {
	ui, err := strconv.ParseUint(string(c), 2, chunkSize)
	if err != nil {
		log.Fatal("Cant parse binary chunk")
	}

	res := strings.ToUpper(fmt.Sprintf("%x", ui))

	return HexChunk(res)
}
