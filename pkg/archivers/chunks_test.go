package chunks

import (
	"reflect"
	"testing"
)

func Test_toBinaryChunks(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "default",
			args: args{
				text: "010101010101110111110101",
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
				text: "010101010101110111",
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
			if got := SplitByChunks(tt.args.text); !reflect.DeepEqual(got, tt.want) {
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
			if got := tt.cs.ToHex(); !reflect.DeepEqual(got, tt.want) {
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

func TestNewHexChunks(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want HexChunks
	}{
		{
			name: "default",
			args: args{
				s: "AB AC 14 D0",
			},
			want: HexChunks{
				HexChunk("AB"),
				HexChunk("AC"),
				HexChunk("14"),
				HexChunk("D0"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		h    HexChunks
		want BinaryChunks
	}{
		{
			name: "default",
			h: HexChunks{
				HexChunk("55"),
				HexChunk("5D"),
				HexChunk("D0"),
			},
			want: BinaryChunks{
				BinaryChunk("01010101"),
				BinaryChunk("01011101"),
				BinaryChunk("11010000"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_String(t *testing.T) {
	tests := []struct {
		name string
		cs   BinaryChunks
		want string
	}{
		{
			name: "default",
			cs: BinaryChunks{
				BinaryChunk("01010101"),
				BinaryChunk("11010101"),
				BinaryChunk("01010101"),
				BinaryChunk("01010000"),
			},
			want: "01010101110101010101010101010000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
