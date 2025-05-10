package vlc

import (
	chunks "archiver/pkg/archivers"
	"archiver/pkg/archivers/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

type Encoder struct {
	Ext       string
	generator table.Generator
}

func NewEncoder(gen table.Generator) Encoder {
	return Encoder{Ext: ".vlc", generator: gen}
}

func (e Encoder) GetExt() string {
	return e.Ext
}

func (e Encoder) Encode(p string) []byte {
	encodingTable := e.generator.NewTable(p)
	return buildEncodedFile(encodingTable, p)
}

func buildEncodedFile(tbl table.EncodingTable, p string) []byte {
	serializedTable := serializeTable(tbl)
	chunkBytes := chunks.SplitByChunks(convertToBin(p, tbl)).ToBytes()

	var res bytes.Buffer

	res.Write(serializeInt(uint32(len(serializedTable))))
	res.Write(serializedTable)
	res.Write(chunkBytes)

	return res.Bytes()
}

func serializeTable(et table.EncodingTable) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, len(et)))
	gobEnc := gob.NewEncoder(buf)

	err := gobEnc.Encode(et)
	if err != nil {
		log.Fatal("Error serializing table ", err)
	}

	return buf.Bytes()
}

func serializeInt(num uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, num)

	return buf
}

func convertToBin(text string, t table.EncodingTable) string {
	buf := strings.Builder{}

	for _, ch := range text {
		buf.WriteString(bin(ch, t))
	}

	return buf.String()
}

func bin(ch rune, t table.EncodingTable) string {
	if v, ok := t[ch]; ok {
		return v
	}

	log.Fatal(fmt.Sprintf("Provided character not found in encoding table = %d", ch))
	return ""
}
