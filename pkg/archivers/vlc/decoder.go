package vlc

import (
	chunks "archiver/pkg/archivers"
	"archiver/pkg/archivers/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
)

type Decoder struct {
	Ext string
	gen table.Generator
}

type DecodedTextInfo struct {
	textLen uint32
}

func NewDecoder() Decoder {
	return Decoder{Ext: ".txt"}
}

func (d Decoder) Decode(p []byte) string {
	const headerBytes = 4

	tableLen, data := binary.BigEndian.Uint32(p[:headerBytes]), p[headerBytes:]
	binaryTable := data[:tableLen]
	binaryText := data[tableLen:]

	decTree := deserializeTable(binaryTable)
	binChunks := chunks.NewBinChunks(binaryText).String()
	decodedText := decTree.Decode(binChunks)

	return decodedText
}

func (d Decoder) GetExt() string {
	return d.Ext
}

func deserializeTable(tableBytes []byte) table.EncodingTable {
	buf := bytes.NewReader(tableBytes)
	gobDec := gob.NewDecoder(buf)

	var t table.EncodingTable

	err := gobDec.Decode(&t)
	if err != nil {
		log.Fatal("Error decoding table", err)
	}

	return t
}
