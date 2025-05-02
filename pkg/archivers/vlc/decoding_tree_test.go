package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "default",
			et: encodingTable{
				' ': "11",
				't': "1001",
				's': "0101",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Right: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: 's',
							},
						},
					},
				},
				Right: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: 't',
							},
						},
					},
					Right: &DecodingTree{
						Value: ' ',
					},
				},
			},
		},
		{
			name: "simple",
			et: encodingTable{
				' ': "11",
			},
			want: DecodingTree{
				Right: &DecodingTree{
					Right: &DecodingTree{
						Value: ' ',
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodingTree_Decode(t *testing.T) {
	type fields struct {
		Value rune
		Left  *DecodingTree
		Right *DecodingTree
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "simple",
			fields: fields{
				Left: &DecodingTree{
					Right: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: 's',
							},
						},
					},
				},
				Right: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: 't',
							},
						},
					},
					Right: &DecodingTree{
						Value: ' ',
					},
				},
			},
			args: args{
				code: "1110010101000000",
			},
			want: " ts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DecodingTree{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := dt.Decode(tt.args.code); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
