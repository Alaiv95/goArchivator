package vlc

import (
	chunks "archiver/pkg/archivers"
	"strings"
	"unicode"
)

func Decode(p []byte) string {
	binChunks := chunks.NewBinChunks(p)
	decTree := getEncodingTable().DecodingTree()
	decodedText := decTree.Decode(binChunks.String())

	return prepDecodedText(decodedText)
}

func prepDecodedText(p string) string {
	buf := strings.Builder{}
	runes := []rune(p)
	iter := 0

	for iter < len(runes) {
		r := runes[iter]

		if runes[iter] == '!' {
			iter++
			r = unicode.ToUpper(runes[iter])
		}

		buf.WriteRune(r)
		iter++
	}

	return buf.String()
}
