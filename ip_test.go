package ipnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIP(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    IP
		wantErr bool
	}{
		{
			"IPv4_Test1",
			args{
				"10.0.0.1",
			},
			v4IP{address: address{0x0, 0xa000001}},
			false,
		},
		{
			"IPv4_Test2",
			args{
				"10.0.0.",
			},
			nil,
			true,
		},
		{
			"IPv6_Test1",
			args{
				"2605:6440:3008:9000::64cc",
			},
			v6IP{address: address{0x2605644030089000, 0x64cc}},
			false,
		},
		{
			"IPv6_Test2",
			args{
				"2605:6440:3008:9000::64cz",
			},
			nil,
			true,
		},
		{
			"UnknownHex",
			args{
				"abcdef",
			},
			nil,
			true,
		},
		{
			"UnknownDec",
			args{
				"1928394",
			},
			nil,
			true,
		},
		{
			"Unknown",
			args{
				"zxy123",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIP(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Want %v, got %v", tt.wantErr, err)
			}
			assert.Equalf(t, tt.want, got, "ParseIP(%v)", tt.args.addr)
		})
	}
}
