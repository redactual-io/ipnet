package ipnet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePrefix(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    Prefix
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"IPv4_Test1",
			args{
				"10.0.10.0/24",
			},
			&v4Prefix{
				first:       address{0x0, 0xa000a00},
				cursor:      address{0x0, 0xa000a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa000aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			assert.NoError,
		},
		{
			"IPv4_Test2",
			args{
				"10.0.10./24",
			},
			nil,
			assert.Error,
		},
		{
			"IPv6_Test1",
			args{
				"2600:c46:0:34::/64",
			},
			&v6Prefix{
				first:       address{0x26000c4600000034, 0x0},
				cursor:      address{0x26000c4600000034, 0x0},
				inverseMask: address{0x0, 0xffffffffffffffff},
				last:        address{0x26000c4600000034, 0xffffffffffffffff},
				len:         64,
				mask:        address{0xffffffffffffffff, 0x0},
			},
			assert.NoError,
		},
		{
			"IPv6_Test2",
			args{
				"2600:c46:z:34::/64",
			},
			nil,
			assert.Error,
		},
		{
			"Unknown",
			args{
				"2600z",
			},
			nil,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePrefix(tt.args.p)
			if !tt.wantErr(t, err, fmt.Sprintf("ParsePrefix(%v)", tt.args.p)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParsePrefix(%v)", tt.args.p)
		})
	}
}
