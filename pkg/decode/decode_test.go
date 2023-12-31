package decode

import (
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"os"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	happy, _ := os.ReadFile("testfiles/happy.data")
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			name:    "happy",
			args:    args{data: decrypt.Decrypt(happy)},
			wantErr: false,
			want: Data{
				DataLoggerSerial: "XGD5BF42MX",
				Serial:           "WRG0BH70JU",
				PowerIn:          295,
				PV1Voltage:       180.7,
				PV1Current:       0.8,
				PV1Power:         154.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
