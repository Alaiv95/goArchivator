package chunks

type Encoder interface {
	Encode(p string) []byte
	GetExt() string
}
type Decoder interface {
	Decode(p []byte) string
	GetExt() string
}
