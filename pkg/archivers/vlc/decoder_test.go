package vlc

import "testing"

func Test_prepDecodedText(t *testing.T) {
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
			args: args{
				p: "!mdu!ds",
			},
			want: "MduDs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepDecodedText(tt.args.p); got != tt.want {
				t.Errorf("prepDecodedText() = %v, want %v", got, tt.want)
			}
		})
	}
}
