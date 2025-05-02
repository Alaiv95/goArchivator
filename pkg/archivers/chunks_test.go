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

func TestNewBinChunks(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "default",
			args: args{
				p: []byte{20, 30, 60, 18},
			},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		cs   BinaryChunks
		want []byte
	}{
		{
			name: "default",
			cs:   BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
			want: []byte{20, 30, 60, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.ToBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
