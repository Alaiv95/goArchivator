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

const ChunkSize = 8
const hexChunkSep = " "

func (h HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(h))

	for _, chunk := range h {
		res = append(res, chunk.ToBinary())
	}

	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {
	const op = "chunks.ToBinary"

	num, err := strconv.ParseUint(string(hc), 16, ChunkSize)
	if err != nil {
		log.Fatal(fmt.Errorf("%s: %w", op, err))
	}

	res := strings.ToUpper(fmt.Sprintf("%b", num))

	if len(res) < ChunkSize {
		prefix := strings.Repeat("0", ChunkSize-len(res))
		res = prefix + res
	}

	return BinaryChunk(res)
}

func NewHexChunks(s string) HexChunks {
	sp := strings.Split(s, hexChunkSep)
	res := make(HexChunks, 0, len(sp))

	for _, c := range sp {
		res = append(res, HexChunk(c))
	}

	return res
}

func SplitByChunks(text string) BinaryChunks {
	strlen := utf8.RuneCountInString(text)
	chunksCnt := strlen / ChunkSize
	res := make(BinaryChunks, 0, chunksCnt)

	buf := strings.Builder{}

	if strlen%ChunkSize != 0 {
		chunksCnt++
	}

	for i, v := range text {
		buf.WriteRune(v)

		if (i+1)%ChunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", ChunkSize-len(lastChunk))
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

func (cs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(cs))

	for _, c := range cs {
		res = append(res, c.toHex())
	}

	return res
}

func (c BinaryChunk) toHex() HexChunk {
	ui, err := strconv.ParseUint(string(c), 2, ChunkSize)
	if err != nil {
		log.Fatal("Cant parse binary chunk")
	}

	res := strings.ToUpper(fmt.Sprintf("%x", ui))

	return HexChunk(res)
}

func (cs BinaryChunks) String() string {
	buf := strings.Builder{}

	for _, v := range cs {
		buf.WriteString(string(v))
	}

	return buf.String()
}
