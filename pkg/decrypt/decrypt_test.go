package decrypt

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "happy path",
			args: args{data: []byte("\x00\x02\x00\x06\x02\x41\x01\x03\x1f\x35\x2b\x42\x23\x32\x40\x75\x3f\x37")},
			want: []byte("\x00\x02\x00\x06\x02\x41\x01\x03\x58\x47\x44\x35\x42\x46\x34\x32\x4d\x58"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
