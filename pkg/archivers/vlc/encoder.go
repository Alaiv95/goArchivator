package vlc

import (
	chunks "archiver/pkg/archivers"
	"log"
	"strings"
	"unicode"
)

type Encoder struct {
	Ext string
}

func NewEncoder() Encoder {
	return Encoder{Ext: ".vlc"}
}

func (e Encoder) GetExt() string {
	return e.Ext
}

func (e Encoder) Encode(p string) []byte {
	binaryText := convertToBin(prepText(p))
	binChunks := chunks.SplitByChunks(binaryText)

	return binChunks.ToBytes()
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
