package vlc

import (
	chunks "archiver/pkg/archivers"
	"log"
	"strings"
	"unicode"
)

type encodingTable map[rune]string

const chunkSize = 8

func Encode(p string) string {
	binaryText := convertToBin(prepText(p))
	binChunks := chunks.ToBinaryChunks(binaryText, chunkSize)

	return binChunks.ToHex(chunkSize).ToString()
}

func convertToBin(text string) string {
	buf := strings.Builder{}

	for _, ch := range text {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()
	if v, ok := table[ch]; ok {
		return v
	}

	log.Fatal("Provided character not found in encoding table")
	return ""
}

func prepText(p string) string {
	builder := strings.Builder{}

	for _, v := range p {
		if unicode.IsUpper(v) {
			builder.WriteRune('!')
			builder.WriteRune(unicode.ToLower(v))
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
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
