package chunks

import (
	"reflect"
	"testing"
)

func Test_toBinaryChunks(t *testing.T) {
	type args struct {
		text      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "default",
			args: args{
				text:      "010101010101110111110101",
				chunkSize: 8,
			},
			want: BinaryChunks{
				BinaryChunk("01010101"),
				BinaryChunk("01011101"),
				BinaryChunk("11110101"),
			},
		},
		{
			name: "last chunks less then size",
			args: args{
				text:      "010101010101110111",
				chunkSize: 8,
			},
			want: BinaryChunks{
				BinaryChunk("01010101"),
				BinaryChunk("01011101"),
				BinaryChunk("11000000"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBinaryChunks(tt.args.text, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toBinaryChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_toHex(t *testing.T) {
	tests := []struct {
		name string
		cs   BinaryChunks
		want HexChunks
	}{
		{
			name: "default",
			cs: BinaryChunks{
				BinaryChunk("01010101"),
				BinaryChunk("01011101"),
				BinaryChunk("11010000"),
			},
			want: HexChunks{
				HexChunk("55"),
				HexChunk("5D"),
				HexChunk("D0"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.ToHex(8); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_toString(t *testing.T) {
	tests := []struct {
		name string
		h    HexChunks
		want string
	}{
		{
			name: "default",
			h: HexChunks{
				HexChunk("55"),
				HexChunk("5D"),
				HexChunk("50"),
			},
			want: "55 5D 50",
		},
		{
			name: "one len",
			h: HexChunks{
				HexChunk("55"),
			},
			want: "55",
		},
		{
			name: "empty",
			h:    HexChunks{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ToString(); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
