package shannon_fano

import (
	"archiver/pkg/archivers/vlc/table"
	"reflect"
	"testing"
)

func Test_bestDividerPos(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				codes: []code{
					{
						Value: 'a',
						Count: 4,
					},
					{
						Value: 'b',
						Count: 2,
					},
					{
						Value: 'c',
						Count: 1,
					},
					{
						Value: 'd',
						Count: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "default",
			args: args{
				codes: []code{
					{
						Value: 'a',
						Count: 4,
					},
					{
						Value: 'v',
						Count: 1,
					},
					{
						Value: 'b',
						Count: 2,
					},
					{
						Value: 'c',
						Count: 2,
					},
				},
			},
			want: 1,
		},
		{
			name: "default",
			args: args{
				codes: []code{
					{
						Value: 'a',
						Count: 6,
					},
					{
						Value: 'v',
						Count: 1,
					},
					{
						Value: 'b',
						Count: 2,
					},
					{
						Value: 'c',
						Count: 2,
					},
				},
			},
			want: 1,
		},
		{
			name: "default",
			args: args{
				codes: []code{
					{
						Value: 'a',
						Count: 6,
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPos(tt.args.codes); got != tt.want {
				t.Errorf("bestDividerPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignBits(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
		want []code
	}{
		{
			name: "default",
			args: args{
				codes: []code{
					{
						Value: 'a',
						Count: 4,
					},
					{
						Value: 'b',
						Count: 2,
					},
					{
						Value: 'c',
						Count: 1,
					},
					{
						Value: 'd',
						Count: 1,
					},
				},
			},
			want: []code{
				{
					Value: 'a',
					Count: 4,
					Size:  1,
					Bit:   0,
				},
				{
					Value: 'b',
					Count: 2,
					Size:  2,
					Bit:   2,
				},
				{
					Value: 'c',
					Count: 1,
					Bit:   6,
					Size:  3,
				},
				{
					Value: 'd',
					Count: 1,
					Bit:   7,
					Size:  3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignBits(tt.args.codes)
			if !reflect.DeepEqual(tt.args.codes, tt.want) {
				t.Errorf("assignBits() = %v, want %v", tt.args.codes, tt.want)
			}
		})
	}
}

func Test_buildTable(t *testing.T) {
	type args struct {
		occurrence charOccurrence
	}
	tests := []struct {
		name string
		args args
		want encodingTable
	}{
		{
			name: "default",
			args: args{
				occurrence: charOccurrence{
					'a': 4,
					'b': 2,
					'c': 1,
					'd': 1,
				},
			},
			want: encodingTable{
				'a': code{
					Value: 'a',
					Count: 4,
					Size:  1,
					Bit:   0,
				},
				'b': code{Value: 'b',
					Count: 2,
					Size:  2,
					Bit:   2},
				'c': code{Value: 'c',
					Count: 1,
					Bit:   6,
					Size:  3},
				'd': code{Value: 'd',
					Count: 1,
					Bit:   7,
					Size:  3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTable(tt.args.occurrence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodingTable_export(t *testing.T) {
	tests := []struct {
		name string
		t    encodingTable
		want table.EncodingTable
	}{
		{
			name: "default",
			t: encodingTable{
				'a': code{
					Value: 'a',
					Count: 4,
					Size:  1,
					Bit:   0,
				},
				'b': code{Value: 'b',
					Count: 2,
					Size:  2,
					Bit:   2},
				'c': code{Value: 'c',
					Count: 1,
					Bit:   6,
					Size:  3},
				'd': code{Value: 'd',
					Count: 1,
					Bit:   7,
					Size:  3},
			},
			want: table.EncodingTable{
				'a': "0",
				'b': "10",
				'c': "110",
				'd': "111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.export(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("export() = %v, want %v", got, tt.want)
			}
		})
	}
}
