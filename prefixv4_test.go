package ipnet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ipv4Mask(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name            string
		args            args
		wantMask        address
		wantInverseMask address
	}{
		{"Test0", args{0}, address{0x0, 0x0}, address{0x0, 0xffffffff}},
		{"Test1", args{1}, address{0x0, 0x80000000}, address{0x0, 0x7fffffff}},
		{"Test2", args{2}, address{0x0, 0xc0000000}, address{0x0, 0x3fffffff}},
		{"Test3", args{3}, address{0x0, 0xe0000000}, address{0x0, 0x1fffffff}},
		{"Test4", args{4}, address{0x0, 0xf0000000}, address{0x0, 0xfffffff}},
		{"Test5", args{5}, address{0x0, 0xf8000000}, address{0x0, 0x7ffffff}},
		{"Test6", args{6}, address{0x0, 0xfc000000}, address{0x0, 0x3ffffff}},
		{"Test7", args{7}, address{0x0, 0xfe000000}, address{0x0, 0x1ffffff}},
		{"Test8", args{8}, address{0x0, 0xff000000}, address{0x0, 0xffffff}},
		{"Test9", args{9}, address{0x0, 0xff800000}, address{0x0, 0x7fffff}},
		{"Test10", args{10}, address{0x0, 0xffc00000}, address{0x0, 0x3fffff}},
		{"Test11", args{11}, address{0x0, 0xffe00000}, address{0x0, 0x1fffff}},
		{"Test12", args{12}, address{0x0, 0xfff00000}, address{0x0, 0xfffff}},
		{"Test13", args{13}, address{0x0, 0xfff80000}, address{0x0, 0x7ffff}},
		{"Test14", args{14}, address{0x0, 0xfffc0000}, address{0x0, 0x3ffff}},
		{"Test15", args{15}, address{0x0, 0xfffe0000}, address{0x0, 0x1ffff}},
		{"Test16", args{16}, address{0x0, 0xffff0000}, address{0x0, 0xffff}},
		{"Test17", args{17}, address{0x0, 0xffff8000}, address{0x0, 0x7fff}},
		{"Test18", args{18}, address{0x0, 0xffffc000}, address{0x0, 0x3fff}},
		{"Test19", args{19}, address{0x0, 0xffffe000}, address{0x0, 0x1fff}},
		{"Test20", args{20}, address{0x0, 0xfffff000}, address{0x0, 0xfff}},
		{"Test21", args{21}, address{0x0, 0xfffff800}, address{0x0, 0x7ff}},
		{"Test22", args{22}, address{0x0, 0xfffffc00}, address{0x0, 0x3ff}},
		{"Test23", args{23}, address{0x0, 0xfffffe00}, address{0x0, 0x1ff}},
		{"Test24", args{24}, address{0x0, 0xffffff00}, address{0x0, 0xff}},
		{"Test25", args{25}, address{0x0, 0xffffff80}, address{0x0, 0x7f}},
		{"Test26", args{26}, address{0x0, 0xffffffc0}, address{0x0, 0x3f}},
		{"Test27", args{27}, address{0x0, 0xffffffe0}, address{0x0, 0x1f}},
		{"Test28", args{28}, address{0x0, 0xfffffff0}, address{0x0, 0xf}},
		{"Test29", args{29}, address{0x0, 0xfffffff8}, address{0x0, 0x7}},
		{"Test30", args{30}, address{0x0, 0xfffffffc}, address{0x0, 0x3}},
		{"Test31", args{31}, address{0x0, 0xfffffffe}, address{0x0, 0x1}},
		{"Test32", args{32}, address{0x0, 0xffffffff}, address{0x0, 0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMask, gotInverseMask := ipv4Mask(tt.args.length)
			assert.Equalf(t, tt.wantMask, gotMask, "ipv4Mask(%v)", tt.args.length)
			assert.Equalf(t, tt.wantInverseMask, gotInverseMask, "ipv4Mask(%v)", tt.args.length)
		})
	}
}

func Test_parseIPv4Length(t *testing.T) {
	type args struct {
		length string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantOk bool
	}{
		{"Test-1", args{"-1"}, 0, false},
		{"Test1", args{"1"}, 1, true},
		{"Test2", args{"2"}, 2, true},
		{"Test3", args{"3"}, 3, true},
		{"Test4", args{"4"}, 4, true},
		{"Test5", args{"5"}, 5, true},
		{"Test6", args{"6"}, 6, true},
		{"Test7", args{"7"}, 7, true},
		{"Test8", args{"8"}, 8, true},
		{"Test9", args{"9"}, 9, true},
		{"Test10", args{"10"}, 10, true},
		{"Test11", args{"11"}, 11, true},
		{"Test12", args{"12"}, 12, true},
		{"Test13", args{"13"}, 13, true},
		{"Test14", args{"14"}, 14, true},
		{"Test15", args{"15"}, 15, true},
		{"Test16", args{"16"}, 16, true},
		{"Test17", args{"17"}, 17, true},
		{"Test18", args{"18"}, 18, true},
		{"Test19", args{"19"}, 19, true},
		{"Test20", args{"20"}, 20, true},
		{"Test21", args{"21"}, 21, true},
		{"Test22", args{"22"}, 22, true},
		{"Test23", args{"23"}, 23, true},
		{"Test24", args{"24"}, 24, true},
		{"Test25", args{"25"}, 25, true},
		{"Test26", args{"26"}, 26, true},
		{"Test27", args{"27"}, 27, true},
		{"Test28", args{"28"}, 28, true},
		{"Test29", args{"29"}, 29, true},
		{"Test30", args{"30"}, 30, true},
		{"Test31", args{"31"}, 31, true},
		{"Test32", args{"32"}, 32, true},
		{"Test33", args{"33"}, 0, false},
		{"Test34", args{"z"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseIPv4Length(tt.args.length)
			assert.Equalf(t, tt.want, got, "parseIPv4Length(%v)", tt.args.length)
			assert.Equalf(t, tt.wantOk, got1, "parseIPv4Length(%v)", tt.args.length)
		})
	}
}

func Test_parseIPv4Prefix(t *testing.T) {
	type args struct {
		pfx string
	}
	tests := []struct {
		name    string
		args    args
		want    *v4Prefix
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Test1",
			args{"10.10.10.0/24"},
			&v4Prefix{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			assert.NoError,
		},
		{
			"Test2",
			args{"10.10.10.0.24"},
			nil,
			assert.Error,
		},
		{
			"Test3",
			args{"10.10.10.0/33"},
			nil,
			assert.Error,
		},
		{
			"Test4",
			args{"10.10.0/32"},
			nil,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIPv4Prefix(tt.args.pfx)
			if !tt.wantErr(t, err, fmt.Sprintf("parseIPv4Prefix(%v)", tt.args.pfx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseIPv4Prefix(%v)", tt.args.pfx)
		})
	}
}

func Test_v4Prefix_Cursor(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0a00}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Cursor(), "Cursor()")
		})
	}
}

func Test_v4Prefix_First(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0a00}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.First(), "First()")
		})
	}
}

func Test_v4Prefix_IsInterface(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			false,
		},
		{
			"Test2",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a01},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.IsInterface(), "IsInterface()")
		})
	}
}

func Test_v4Prefix_Last(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0aff}}),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Last(), "Last()")
		})
	}
}

func Test_v4Prefix_Length(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			24,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Length(), "Length()")
		})
	}
}

func Test_v4Prefix_Mask(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xffffff00}}),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Mask(), "Mask()")
		})
	}
}

func Test_v4Prefix_Member(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	type args struct {
		in IP
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			args{
				IP(v4IP{address{0x0, 0xa0a0a0f}}),
			},
			true,
		},
		{
			"Test2",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			args{
				IP(v4IP{address{0x0, 0xa0a0b01}}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Member(tt.args.in), "Member(%v)", tt.args.in)
		})
	}
}

func Test_v4Prefix_Next(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
		wantOk bool
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0a01}}),
			true,
		},
		{
			"Test2",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0aff},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0aff}}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			got, got1 := p.Next()
			assert.Equalf(t, tt.want, got, "Next()")
			assert.Equalf(t, tt.wantOk, got1, "Next()")
		})
	}
}

func Test_v4Prefix_Prev(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   IP
		want1  bool
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0a00}}),
			false,
		},
		{
			"Test2",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a01},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0xa0a0a00}}),
			true,
		},
		{
			"Test3",
			fields{
				first:       address{0x0, 0x0},
				cursor:      address{0x0, 0x0},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			IP(v4IP{address{0x0, 0x0}}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			got, got1 := p.Prev()
			assert.Equalf(t, tt.want, got, "Prev()")
			assert.Equalf(t, tt.want1, got1, "Prev()")
		})
	}
}

func Test_v4Prefix_Random(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a01},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			var randomIP IP
			assert.NotPanics(t, func() { randomIP = p.Random() }, "Random()")
			assert.True(t, p.Member(randomIP), "Random()")
		})
	}
}

func Test_v4Prefix_String(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			"10.10.10.0/24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.String(), "String()")
		})
	}
}

func Test_v4Prefix_Version(t *testing.T) {
	type fields struct {
		first       address
		cursor      address
		inverseMask address
		last        address
		len         int
		mask        address
	}
	tests := []struct {
		name   string
		fields fields
		want   Ver
	}{
		{
			"Test1",
			fields{
				first:       address{0x0, 0xa0a0a00},
				cursor:      address{0x0, 0xa0a0a00},
				inverseMask: address{0x0, 0xff},
				last:        address{0x0, 0xa0a0aff},
				len:         24,
				mask:        address{0x0, 0xffffff00},
			},
			v4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v4Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			assert.Equalf(t, tt.want, p.Version(), "Version()")
		})
	}
}

func Test_v4SplitLen(t *testing.T) {
	type args struct {
		pfx string
	}
	tests := []struct {
		name   string
		args   args
		want   [2]string
		wantOk bool
	}{
		{
			"Test1",
			args{
				"10.0.0.0/16",
			},
			[2]string{"10.0.0.0", "16"},
			true,
		},
		{
			"Test1",
			args{
				"10.0.0.0.16",
			},
			[2]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := v4SplitLen(tt.args.pfx)
			assert.Equalf(t, tt.want, got, "v4SplitLen(%v)", tt.args.pfx)
			assert.Equalf(t, tt.wantOk, got1, "v4SplitLen(%v)", tt.args.pfx)
		})
	}
}
