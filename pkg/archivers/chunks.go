package chunks

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk
type BinaryChunk string

type HexChunks []HexChunk

type HexChunk string

const ChunkSize = 8

func NewBinChunks(p []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(p))

	for _, b := range p {
		c := fmt.Sprintf("%08b", b)
		res = append(res, BinaryChunk(c))
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

func (cs BinaryChunks) String() string {
	buf := strings.Builder{}

	for _, v := range cs {
		buf.WriteString(string(v))
	}

	return buf.String()
}

func (cs BinaryChunks) ToBytes() []byte {
	res := make([]byte, 0, len(cs))

	for _, v := range cs {
		res = append(res, v.Byte())
	}

	return res
}

func (c BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(c), 2, ChunkSize)
	if err != nil {
		panic(err)
	}

	return byte(num)
}
