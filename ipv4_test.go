package ipnet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_v4IP_addr(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   address
	}{
		{
			"Test1",
			fields{
				address{0x0, 0x0a0a0a0a},
			},
			address{0x0, 0x0a0a0a0a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v4IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.addr(), "addr()")
		})
	}
}

func Test_v4IP_String(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test1",
			fields{
				address{0x0, 0x0a0a0a0a},
			},
			"10.10.10.10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v4IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.String(), "String()")
		})
	}
}

func Test_v4IP_Version(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   Ver
	}{
		{
			"Test1",
			fields{
				address{0x0, 0x0a0a0a0a},
			},
			v4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v4IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.Version(), "Version()")
		})
	}
}

func Test_splitV4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		wantOk bool
	}{
		{
			"Test1",
			args{
				"1.2.3.4",
			},
			[]string{"1", "2", "3", "4"},
			true,
		},
		{
			"Test2",
			args{
				"",
			},
			[]string{},
			false,
		},
		{
			"Test3",
			args{
				".2.3.4",
			},
			[]string{},
			false,
		},
		{
			"Test4",
			args{
				"1111.2.3.4",
			},
			[]string{},
			false,
		},
		{
			"Test5",
			args{
				"1.2.3.",
			},
			[]string{},
			false,
		},
		{
			"Test6",
			args{
				"1.2.3.4444",
			},
			[]string{},
			false,
		},
		{
			"Test7",
			args{
				"2.3.4",
			},
			[]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitV4(tt.args.s)
			assert.Equalf(t, tt.want, got, "splitV4(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "splitV4(%v)", tt.args.s)
		})
	}
}

func Test_parseV4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		want   [4]uint64
		wantOk bool
	}{
		{
			"Test1",
			args{
				"10.10.10.10",
			},
			[4]uint64{0xa, 0xa, 0xa, 0xa},
			true,
		},
		{
			"Test2",
			args{
				"10.10.10",
			},
			[4]uint64{0x0, 0x0, 0x0, 0x0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseV4(tt.args.s)
			assert.Equalf(t, tt.want, got, "parseV4(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "parseV4(%v)", tt.args.s)
		})
	}
}

func Test_a10toUint64(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name   string
		args   args
		want   [4]uint64
		wantOk bool
	}{
		{
			"Test1",
			args{
				[]string{"10", "10", "10", "10"},
			},
			[4]uint64{10, 10, 10, 10},
			true,
		},
		{
			"Test2",
			args{
				[]string{},
			},
			[4]uint64{},
			false,
		},
		{
			"Test3",
			args{
				[]string{"10", "10", "10", "1a"},
			},
			[4]uint64{},
			false,
		},
		{
			"Test4",
			args{
				[]string{"10", "256", "10", "10"},
			},
			[4]uint64{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := a10toUint64(tt.args.s)
			assert.Equalf(t, tt.want, got, "a10toUint64(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "a10toUint64(%v)", tt.args.s)
		})
	}
}

func Test_parseIPv4(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    v4IP
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Test1",
			args{
				"10.10.10.10",
			},
			v4IP{address: address{0x0, 0xa0a0a0a}},
			assert.NoError,
		},
		{
			"Test2",
			args{
				"0.0.0.",
			},
			v4IP{},
			assert.Error,
		},
		{
			"Test3",
			args{
				"1.2.3.a",
			},
			v4IP{},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIPv4(tt.args.addr)
			if !tt.wantErr(t, err, fmt.Sprintf("parseIPv4(%v)", tt.args.addr)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseIPv4(%v)", tt.args.addr)
		})
	}
}
