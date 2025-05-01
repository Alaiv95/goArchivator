package vlc

import (
	"testing"
)

func Test_prepText(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{p: "Big"},
			want: "!big",
		},
		{
			name: "empty",
			args: args{p: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepText(tt.args.p); got != tt.want {
				t.Errorf("prepText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToBin(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{text: "!big"},
			want: "0010000000010010010000100",
		},
		{
			name: "empty",
			args: args{text: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBin(tt.args.text); got != tt.want {
				t.Errorf("convertToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
