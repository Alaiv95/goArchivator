package vlc

import (
	"archiver/pkg/archivers/vlc/table"
	"testing"
)

func Test_convertToBin(t *testing.T) {
	type args struct {
		text  string
		table table.EncodingTable
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{text: "!big", table: table.EncodingTable{
				'!': "00",
				'b': "01",
				'i': "10",
				'g': "11",
			}},
			want: "00011011",
		},
		{
			name: "empty",
			args: args{text: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBin(tt.args.text, tt.args.table); got != tt.want {
				t.Errorf("convertToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
