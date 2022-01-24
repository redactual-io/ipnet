package ipnet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_v6Prefix_Member(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			args{
				IP(v6IP{address{0x2a05d07aa0000000, 0xcd85d07}}),
			},
			true,
		},
		{
			"Test2",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			args{ //2400:6500:0:7800::
				IP(v6IP{address{0x2400650007800000, 0xcd85d07}}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Version(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			v6,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_First(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0000000, 0x0}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Last(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0ffffff, 0xffffffffffffffff}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Mask(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0xffffffffff000000, 0x0}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Length(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Cursor(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0000000, 0x0}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_Next(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0000000, 0x1}}),
			true,
		},
		{
			"Test2",
			fields{
				cursor:      address{0xffffffffffffffff, 0xffffffffffffffff},
				first:       address{0xffffffffffffffff, 0x0},
				last:        address{0xffffffffffffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffffffffff, 0x0},
				inverseMask: address{0x0, 0xffffffffffffffff},
				len:         64,
			},
			IP(v6IP{address{0xffffffffffffffff, 0xffffffffffffffff}}),
			false,
		},
		{
			"Test3",
			fields{
				cursor:      address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0ffffff, 0xffffffffffffffff}}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &v6Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			got, got1 := p.Next()
			assert.Equalf(t, tt.want, got, "Next()")
			assert.Equalf(t, tt.want1, got1, "Next()")
		})
	}
}

func Test_v6Prefix_Prev(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x1},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0000000, 0x0}}),
			true,
		},
		{
			"Test2",
			fields{
				cursor:      address{0x0000000000000000, 0x0},
				first:       address{0x0000000000000000, 0x0},
				last:        address{0x0000000000000000, 0xffffffffffffffff},
				mask:        address{0xffffffffffffffff, 0x0},
				inverseMask: address{0x0, 0xffffffffffffffff},
				len:         64,
			},
			IP(v6IP{address{0x0000000000000000, 0x0}}),
			false,
		},
		{
			"Test3",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			IP(v6IP{address{0x2a05d07aa0000000, 0x0}}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &v6Prefix{
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

func Test_v6Prefix_Random(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
		},
		{
			"Test2",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
		},
		{
			"Test3",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
		},
		{
			"Test4",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
		},
		{
			"Test5",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
				first:       tt.fields.first,
				cursor:      tt.fields.cursor,
				inverseMask: tt.fields.inverseMask,
				last:        tt.fields.last,
				len:         tt.fields.len,
				mask:        tt.fields.mask,
			}
			ra := p.Random()
			assert.True(t, p.first.le(ra.addr()))
			assert.True(t, p.last.ge(ra.addr()))
		})
	}
}

func BenchmarkV6Prefix_Random(b *testing.B) {
	b.Run("Random", func(b *testing.B) {
		var ip IP
		pfx := Prefix(
			&v6Prefix{ // "2620:107:4000:5::/64"
				first:       address{0x2620010740000005, 0x0},
				cursor:      address{0x2620010740000005, 0x0},
				inverseMask: address{0x0, 0xffffffffffffffff},
				last:        address{0x2620010740000005, 0xffffffffffffffff},
				len:         64,
				mask:        address{0xffffffffffffffff, 0x0},
			},
		)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ip = pfx.Random()
		}
		b.StopTimer()
		_ = ip.String()
	})
}

func Test_v6Prefix_String(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x1},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			"2a05:d07a:a000:0:0:0:0:0/40",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_v6Prefix_IsInterface(t *testing.T) {
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
				cursor:      address{0x2a05d07aa0000000, 0x1},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			true,
		},
		{
			"Test2",
			fields{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := v6Prefix{
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

func Test_parseIPv6Prefix(t *testing.T) {
	type args struct {
		pfx string
	}
	tests := []struct {
		name    string
		args    args
		want    *v6Prefix
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Test1",
			args{
				"2a05:d07a:a000::/40",
			},
			&v6Prefix{
				cursor:      address{0x2a05d07aa0000000, 0x0},
				first:       address{0x2a05d07aa0000000, 0x0},
				last:        address{0x2a05d07aa0ffffff, 0xffffffffffffffff},
				mask:        address{0xffffffffff000000, 0x0},
				inverseMask: address{0xffffff, 0xffffffffffffffff},
				len:         40,
			},
			assert.NoError,
		},
		{
			"Test2",
			args{
				"2a05:d07a:a000::/129",
			},
			nil,
			assert.Error,
		},
		{
			"Test3",
			args{
				"2a05:d07a:z000::/77",
			},
			nil,
			assert.Error,
		},
		{
			"Test4",
			args{
				"2a05:d07a:40000::/77",
			},
			nil,
			assert.Error,
		},
		{
			"Test5",
			args{
				"2a05:d07a:40000::77",
			},
			nil,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIPv6Prefix(tt.args.pfx)
			if !tt.wantErr(t, err, fmt.Sprintf("parseIPv6Prefix(%v)", tt.args.pfx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseIPv6Prefix(%v)", tt.args.pfx)
		})
	}
}

func Test_ipv6Mask(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name            string
		args            args
		wantMask        address
		wantInverseMask address
	}{
		{"Test0", args{length: 0}, address{0x0, 0x0}, address{0xffffffffffffffff, 0xffffffffffffffff}},
		{"Test1", args{length: 1}, address{0x8000000000000000, 0x0}, address{0x7fffffffffffffff, 0xffffffffffffffff}},
		{"Test2", args{length: 2}, address{0xc000000000000000, 0x0}, address{0x3fffffffffffffff, 0xffffffffffffffff}},
		{"Test3", args{length: 3}, address{0xe000000000000000, 0x0}, address{0x1fffffffffffffff, 0xffffffffffffffff}},
		{"Test4", args{length: 4}, address{0xf000000000000000, 0x0}, address{0xfffffffffffffff, 0xffffffffffffffff}},
		{"Test5", args{length: 5}, address{0xf800000000000000, 0x0}, address{0x7ffffffffffffff, 0xffffffffffffffff}},
		{"Test6", args{length: 6}, address{0xfc00000000000000, 0x0}, address{0x3ffffffffffffff, 0xffffffffffffffff}},
		{"Test7", args{length: 7}, address{0xfe00000000000000, 0x0}, address{0x1ffffffffffffff, 0xffffffffffffffff}},
		{"Test8", args{length: 8}, address{0xff00000000000000, 0x0}, address{0xffffffffffffff, 0xffffffffffffffff}},
		{"Test9", args{length: 9}, address{0xff80000000000000, 0x0}, address{0x7fffffffffffff, 0xffffffffffffffff}},
		{"Test10", args{length: 10}, address{0xffc0000000000000, 0x0}, address{0x3fffffffffffff, 0xffffffffffffffff}},
		{"Test11", args{length: 11}, address{0xffe0000000000000, 0x0}, address{0x1fffffffffffff, 0xffffffffffffffff}},
		{"Test12", args{length: 12}, address{0xfff0000000000000, 0x0}, address{0xfffffffffffff, 0xffffffffffffffff}},
		{"Test13", args{length: 13}, address{0xfff8000000000000, 0x0}, address{0x7ffffffffffff, 0xffffffffffffffff}},
		{"Test14", args{length: 14}, address{0xfffc000000000000, 0x0}, address{0x3ffffffffffff, 0xffffffffffffffff}},
		{"Test15", args{length: 15}, address{0xfffe000000000000, 0x0}, address{0x1ffffffffffff, 0xffffffffffffffff}},
		{"Test16", args{length: 16}, address{0xffff000000000000, 0x0}, address{0xffffffffffff, 0xffffffffffffffff}},
		{"Test17", args{length: 17}, address{0xffff800000000000, 0x0}, address{0x7fffffffffff, 0xffffffffffffffff}},
		{"Test18", args{length: 18}, address{0xffffc00000000000, 0x0}, address{0x3fffffffffff, 0xffffffffffffffff}},
		{"Test19", args{length: 19}, address{0xffffe00000000000, 0x0}, address{0x1fffffffffff, 0xffffffffffffffff}},
		{"Test20", args{length: 20}, address{0xfffff00000000000, 0x0}, address{0xfffffffffff, 0xffffffffffffffff}},
		{"Test21", args{length: 21}, address{0xfffff80000000000, 0x0}, address{0x7ffffffffff, 0xffffffffffffffff}},
		{"Test22", args{length: 22}, address{0xfffffc0000000000, 0x0}, address{0x3ffffffffff, 0xffffffffffffffff}},
		{"Test23", args{length: 23}, address{0xfffffe0000000000, 0x0}, address{0x1ffffffffff, 0xffffffffffffffff}},
		{"Test24", args{length: 24}, address{0xffffff0000000000, 0x0}, address{0xffffffffff, 0xffffffffffffffff}},
		{"Test25", args{length: 25}, address{0xffffff8000000000, 0x0}, address{0x7fffffffff, 0xffffffffffffffff}},
		{"Test26", args{length: 26}, address{0xffffffc000000000, 0x0}, address{0x3fffffffff, 0xffffffffffffffff}},
		{"Test27", args{length: 27}, address{0xffffffe000000000, 0x0}, address{0x1fffffffff, 0xffffffffffffffff}},
		{"Test28", args{length: 28}, address{0xfffffff000000000, 0x0}, address{0xfffffffff, 0xffffffffffffffff}},
		{"Test29", args{length: 29}, address{0xfffffff800000000, 0x0}, address{0x7ffffffff, 0xffffffffffffffff}},
		{"Test30", args{length: 30}, address{0xfffffffc00000000, 0x0}, address{0x3ffffffff, 0xffffffffffffffff}},
		{"Test31", args{length: 31}, address{0xfffffffe00000000, 0x0}, address{0x1ffffffff, 0xffffffffffffffff}},
		{"Test32", args{length: 32}, address{0xffffffff00000000, 0x0}, address{0xffffffff, 0xffffffffffffffff}},
		{"Test33", args{length: 33}, address{0xffffffff80000000, 0x0}, address{0x7fffffff, 0xffffffffffffffff}},
		{"Test34", args{length: 34}, address{0xffffffffc0000000, 0x0}, address{0x3fffffff, 0xffffffffffffffff}},
		{"Test35", args{length: 35}, address{0xffffffffe0000000, 0x0}, address{0x1fffffff, 0xffffffffffffffff}},
		{"Test36", args{length: 36}, address{0xfffffffff0000000, 0x0}, address{0xfffffff, 0xffffffffffffffff}},
		{"Test37", args{length: 37}, address{0xfffffffff8000000, 0x0}, address{0x7ffffff, 0xffffffffffffffff}},
		{"Test38", args{length: 38}, address{0xfffffffffc000000, 0x0}, address{0x3ffffff, 0xffffffffffffffff}},
		{"Test39", args{length: 39}, address{0xfffffffffe000000, 0x0}, address{0x1ffffff, 0xffffffffffffffff}},
		{"Test40", args{length: 40}, address{0xffffffffff000000, 0x0}, address{0xffffff, 0xffffffffffffffff}},
		{"Test41", args{length: 41}, address{0xffffffffff800000, 0x0}, address{0x7fffff, 0xffffffffffffffff}},
		{"Test42", args{length: 42}, address{0xffffffffffc00000, 0x0}, address{0x3fffff, 0xffffffffffffffff}},
		{"Test43", args{length: 43}, address{0xffffffffffe00000, 0x0}, address{0x1fffff, 0xffffffffffffffff}},
		{"Test44", args{length: 44}, address{0xfffffffffff00000, 0x0}, address{0xfffff, 0xffffffffffffffff}},
		{"Test45", args{length: 45}, address{0xfffffffffff80000, 0x0}, address{0x7ffff, 0xffffffffffffffff}},
		{"Test46", args{length: 46}, address{0xfffffffffffc0000, 0x0}, address{0x3ffff, 0xffffffffffffffff}},
		{"Test47", args{length: 47}, address{0xfffffffffffe0000, 0x0}, address{0x1ffff, 0xffffffffffffffff}},
		{"Test48", args{length: 48}, address{0xffffffffffff0000, 0x0}, address{0xffff, 0xffffffffffffffff}},
		{"Test49", args{length: 49}, address{0xffffffffffff8000, 0x0}, address{0x7fff, 0xffffffffffffffff}},
		{"Test50", args{length: 50}, address{0xffffffffffffc000, 0x0}, address{0x3fff, 0xffffffffffffffff}},
		{"Test51", args{length: 51}, address{0xffffffffffffe000, 0x0}, address{0x1fff, 0xffffffffffffffff}},
		{"Test52", args{length: 52}, address{0xfffffffffffff000, 0x0}, address{0xfff, 0xffffffffffffffff}},
		{"Test53", args{length: 53}, address{0xfffffffffffff800, 0x0}, address{0x7ff, 0xffffffffffffffff}},
		{"Test54", args{length: 54}, address{0xfffffffffffffc00, 0x0}, address{0x3ff, 0xffffffffffffffff}},
		{"Test55", args{length: 55}, address{0xfffffffffffffe00, 0x0}, address{0x1ff, 0xffffffffffffffff}},
		{"Test56", args{length: 56}, address{0xffffffffffffff00, 0x0}, address{0xff, 0xffffffffffffffff}},
		{"Test57", args{length: 57}, address{0xffffffffffffff80, 0x0}, address{0x7f, 0xffffffffffffffff}},
		{"Test58", args{length: 58}, address{0xffffffffffffffc0, 0x0}, address{0x3f, 0xffffffffffffffff}},
		{"Test59", args{length: 59}, address{0xffffffffffffffe0, 0x0}, address{0x1f, 0xffffffffffffffff}},
		{"Test60", args{length: 60}, address{0xfffffffffffffff0, 0x0}, address{0xf, 0xffffffffffffffff}},
		{"Test61", args{length: 61}, address{0xfffffffffffffff8, 0x0}, address{0x7, 0xffffffffffffffff}},
		{"Test62", args{length: 62}, address{0xfffffffffffffffc, 0x0}, address{0x3, 0xffffffffffffffff}},
		{"Test63", args{length: 63}, address{0xfffffffffffffffe, 0x0}, address{0x1, 0xffffffffffffffff}},
		{"Test64", args{length: 64}, address{0xffffffffffffffff, 0x0}, address{0x0, 0xffffffffffffffff}},
		{"Test65", args{length: 65}, address{0xffffffffffffffff, 0x8000000000000000}, address{0x0, 0x7fffffffffffffff}},
		{"Test66", args{length: 66}, address{0xffffffffffffffff, 0xc000000000000000}, address{0x0, 0x3fffffffffffffff}},
		{"Test67", args{length: 67}, address{0xffffffffffffffff, 0xe000000000000000}, address{0x0, 0x1fffffffffffffff}},
		{"Test68", args{length: 68}, address{0xffffffffffffffff, 0xf000000000000000}, address{0x0, 0xfffffffffffffff}},
		{"Test69", args{length: 69}, address{0xffffffffffffffff, 0xf800000000000000}, address{0x0, 0x7ffffffffffffff}},
		{"Test70", args{length: 70}, address{0xffffffffffffffff, 0xfc00000000000000}, address{0x0, 0x3ffffffffffffff}},
		{"Test71", args{length: 71}, address{0xffffffffffffffff, 0xfe00000000000000}, address{0x0, 0x1ffffffffffffff}},
		{"Test72", args{length: 72}, address{0xffffffffffffffff, 0xff00000000000000}, address{0x0, 0xffffffffffffff}},
		{"Test73", args{length: 73}, address{0xffffffffffffffff, 0xff80000000000000}, address{0x0, 0x7fffffffffffff}},
		{"Test74", args{length: 74}, address{0xffffffffffffffff, 0xffc0000000000000}, address{0x0, 0x3fffffffffffff}},
		{"Test75", args{length: 75}, address{0xffffffffffffffff, 0xffe0000000000000}, address{0x0, 0x1fffffffffffff}},
		{"Test76", args{length: 76}, address{0xffffffffffffffff, 0xfff0000000000000}, address{0x0, 0xfffffffffffff}},
		{"Test77", args{length: 77}, address{0xffffffffffffffff, 0xfff8000000000000}, address{0x0, 0x7ffffffffffff}},
		{"Test78", args{length: 78}, address{0xffffffffffffffff, 0xfffc000000000000}, address{0x0, 0x3ffffffffffff}},
		{"Test79", args{length: 79}, address{0xffffffffffffffff, 0xfffe000000000000}, address{0x0, 0x1ffffffffffff}},
		{"Test80", args{length: 80}, address{0xffffffffffffffff, 0xffff000000000000}, address{0x0, 0xffffffffffff}},
		{"Test81", args{length: 81}, address{0xffffffffffffffff, 0xffff800000000000}, address{0x0, 0x7fffffffffff}},
		{"Test82", args{length: 82}, address{0xffffffffffffffff, 0xffffc00000000000}, address{0x0, 0x3fffffffffff}},
		{"Test83", args{length: 83}, address{0xffffffffffffffff, 0xffffe00000000000}, address{0x0, 0x1fffffffffff}},
		{"Test84", args{length: 84}, address{0xffffffffffffffff, 0xfffff00000000000}, address{0x0, 0xfffffffffff}},
		{"Test85", args{length: 85}, address{0xffffffffffffffff, 0xfffff80000000000}, address{0x0, 0x7ffffffffff}},
		{"Test86", args{length: 86}, address{0xffffffffffffffff, 0xfffffc0000000000}, address{0x0, 0x3ffffffffff}},
		{"Test87", args{length: 87}, address{0xffffffffffffffff, 0xfffffe0000000000}, address{0x0, 0x1ffffffffff}},
		{"Test88", args{length: 88}, address{0xffffffffffffffff, 0xffffff0000000000}, address{0x0, 0xffffffffff}},
		{"Test89", args{length: 89}, address{0xffffffffffffffff, 0xffffff8000000000}, address{0x0, 0x7fffffffff}},
		{"Test90", args{length: 90}, address{0xffffffffffffffff, 0xffffffc000000000}, address{0x0, 0x3fffffffff}},
		{"Test91", args{length: 91}, address{0xffffffffffffffff, 0xffffffe000000000}, address{0x0, 0x1fffffffff}},
		{"Test92", args{length: 92}, address{0xffffffffffffffff, 0xfffffff000000000}, address{0x0, 0xfffffffff}},
		{"Test93", args{length: 93}, address{0xffffffffffffffff, 0xfffffff800000000}, address{0x0, 0x7ffffffff}},
		{"Test94", args{length: 94}, address{0xffffffffffffffff, 0xfffffffc00000000}, address{0x0, 0x3ffffffff}},
		{"Test95", args{length: 95}, address{0xffffffffffffffff, 0xfffffffe00000000}, address{0x0, 0x1ffffffff}},
		{"Test96", args{length: 96}, address{0xffffffffffffffff, 0xffffffff00000000}, address{0x0, 0xffffffff}},
		{"Test97", args{length: 97}, address{0xffffffffffffffff, 0xffffffff80000000}, address{0x0, 0x7fffffff}},
		{"Test98", args{length: 98}, address{0xffffffffffffffff, 0xffffffffc0000000}, address{0x0, 0x3fffffff}},
		{"Test99", args{length: 99}, address{0xffffffffffffffff, 0xffffffffe0000000}, address{0x0, 0x1fffffff}},
		{"Test100", args{length: 100}, address{0xffffffffffffffff, 0xfffffffff0000000}, address{0x0, 0xfffffff}},
		{"Test101", args{length: 101}, address{0xffffffffffffffff, 0xfffffffff8000000}, address{0x0, 0x7ffffff}},
		{"Test102", args{length: 102}, address{0xffffffffffffffff, 0xfffffffffc000000}, address{0x0, 0x3ffffff}},
		{"Test103", args{length: 103}, address{0xffffffffffffffff, 0xfffffffffe000000}, address{0x0, 0x1ffffff}},
		{"Test104", args{length: 104}, address{0xffffffffffffffff, 0xffffffffff000000}, address{0x0, 0xffffff}},
		{"Test105", args{length: 105}, address{0xffffffffffffffff, 0xffffffffff800000}, address{0x0, 0x7fffff}},
		{"Test106", args{length: 106}, address{0xffffffffffffffff, 0xffffffffffc00000}, address{0x0, 0x3fffff}},
		{"Test107", args{length: 107}, address{0xffffffffffffffff, 0xffffffffffe00000}, address{0x0, 0x1fffff}},
		{"Test108", args{length: 108}, address{0xffffffffffffffff, 0xfffffffffff00000}, address{0x0, 0xfffff}},
		{"Test109", args{length: 109}, address{0xffffffffffffffff, 0xfffffffffff80000}, address{0x0, 0x7ffff}},
		{"Test110", args{length: 110}, address{0xffffffffffffffff, 0xfffffffffffc0000}, address{0x0, 0x3ffff}},
		{"Test111", args{length: 111}, address{0xffffffffffffffff, 0xfffffffffffe0000}, address{0x0, 0x1ffff}},
		{"Test112", args{length: 112}, address{0xffffffffffffffff, 0xffffffffffff0000}, address{0x0, 0xffff}},
		{"Test113", args{length: 113}, address{0xffffffffffffffff, 0xffffffffffff8000}, address{0x0, 0x7fff}},
		{"Test114", args{length: 114}, address{0xffffffffffffffff, 0xffffffffffffc000}, address{0x0, 0x3fff}},
		{"Test115", args{length: 115}, address{0xffffffffffffffff, 0xffffffffffffe000}, address{0x0, 0x1fff}},
		{"Test116", args{length: 116}, address{0xffffffffffffffff, 0xfffffffffffff000}, address{0x0, 0xfff}},
		{"Test117", args{length: 117}, address{0xffffffffffffffff, 0xfffffffffffff800}, address{0x0, 0x7ff}},
		{"Test118", args{length: 118}, address{0xffffffffffffffff, 0xfffffffffffffc00}, address{0x0, 0x3ff}},
		{"Test119", args{length: 119}, address{0xffffffffffffffff, 0xfffffffffffffe00}, address{0x0, 0x1ff}},
		{"Test120", args{length: 120}, address{0xffffffffffffffff, 0xffffffffffffff00}, address{0x0, 0xff}},
		{"Test121", args{length: 121}, address{0xffffffffffffffff, 0xffffffffffffff80}, address{0x0, 0x7f}},
		{"Test122", args{length: 122}, address{0xffffffffffffffff, 0xffffffffffffffc0}, address{0x0, 0x3f}},
		{"Test123", args{length: 123}, address{0xffffffffffffffff, 0xffffffffffffffe0}, address{0x0, 0x1f}},
		{"Test124", args{length: 124}, address{0xffffffffffffffff, 0xfffffffffffffff0}, address{0x0, 0xf}},
		{"Test125", args{length: 125}, address{0xffffffffffffffff, 0xfffffffffffffff8}, address{0x0, 0x7}},
		{"Test126", args{length: 126}, address{0xffffffffffffffff, 0xfffffffffffffffc}, address{0x0, 0x3}},
		{"Test127", args{length: 127}, address{0xffffffffffffffff, 0xfffffffffffffffe}, address{0x0, 0x1}},
		{"Test128", args{length: 128}, address{0xffffffffffffffff, 0xffffffffffffffff}, address{0x0, 0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMask, gotInverseMask := ipv6Mask(tt.args.length)
			assert.Equalf(t, tt.wantMask, gotMask, "ipv6Mask(%v)", tt.args.length)
			assert.Equalf(t, tt.wantInverseMask, gotInverseMask, "ipv6Mask(%v)", tt.args.length)
		})
	}
}

func Test_parseIPv6Length(t *testing.T) {
	type args struct {
		length string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantOk bool
	}{
		{"Test1", args{"-1"}, 0, false},
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
		{"Test33", args{"33"}, 33, true},
		{"Test34", args{"34"}, 34, true},
		{"Test35", args{"35"}, 35, true},
		{"Test36", args{"36"}, 36, true},
		{"Test37", args{"37"}, 37, true},
		{"Test38", args{"38"}, 38, true},
		{"Test39", args{"39"}, 39, true},
		{"Test40", args{"40"}, 40, true},
		{"Test41", args{"41"}, 41, true},
		{"Test42", args{"42"}, 42, true},
		{"Test43", args{"43"}, 43, true},
		{"Test44", args{"44"}, 44, true},
		{"Test45", args{"45"}, 45, true},
		{"Test46", args{"46"}, 46, true},
		{"Test47", args{"47"}, 47, true},
		{"Test48", args{"48"}, 48, true},
		{"Test49", args{"49"}, 49, true},
		{"Test50", args{"50"}, 50, true},
		{"Test51", args{"51"}, 51, true},
		{"Test52", args{"52"}, 52, true},
		{"Test53", args{"53"}, 53, true},
		{"Test54", args{"54"}, 54, true},
		{"Test55", args{"55"}, 55, true},
		{"Test56", args{"56"}, 56, true},
		{"Test57", args{"57"}, 57, true},
		{"Test58", args{"58"}, 58, true},
		{"Test59", args{"59"}, 59, true},
		{"Test60", args{"60"}, 60, true},
		{"Test61", args{"61"}, 61, true},
		{"Test62", args{"62"}, 62, true},
		{"Test63", args{"63"}, 63, true},
		{"Test64", args{"64"}, 64, true},
		{"Test65", args{"65"}, 65, true},
		{"Test66", args{"66"}, 66, true},
		{"Test67", args{"67"}, 67, true},
		{"Test68", args{"68"}, 68, true},
		{"Test69", args{"69"}, 69, true},
		{"Test70", args{"70"}, 70, true},
		{"Test71", args{"71"}, 71, true},
		{"Test72", args{"72"}, 72, true},
		{"Test73", args{"73"}, 73, true},
		{"Test74", args{"74"}, 74, true},
		{"Test75", args{"75"}, 75, true},
		{"Test76", args{"76"}, 76, true},
		{"Test77", args{"77"}, 77, true},
		{"Test78", args{"78"}, 78, true},
		{"Test79", args{"79"}, 79, true},
		{"Test80", args{"80"}, 80, true},
		{"Test81", args{"81"}, 81, true},
		{"Test82", args{"82"}, 82, true},
		{"Test83", args{"83"}, 83, true},
		{"Test84", args{"84"}, 84, true},
		{"Test85", args{"85"}, 85, true},
		{"Test86", args{"86"}, 86, true},
		{"Test87", args{"87"}, 87, true},
		{"Test88", args{"88"}, 88, true},
		{"Test89", args{"89"}, 89, true},
		{"Test90", args{"90"}, 90, true},
		{"Test91", args{"91"}, 91, true},
		{"Test92", args{"92"}, 92, true},
		{"Test93", args{"93"}, 93, true},
		{"Test94", args{"94"}, 94, true},
		{"Test95", args{"95"}, 95, true},
		{"Test96", args{"96"}, 96, true},
		{"Test97", args{"97"}, 97, true},
		{"Test98", args{"98"}, 98, true},
		{"Test99", args{"99"}, 99, true},
		{"Test100", args{"100"}, 100, true},
		{"Test101", args{"101"}, 101, true},
		{"Test102", args{"102"}, 102, true},
		{"Test103", args{"103"}, 103, true},
		{"Test104", args{"104"}, 104, true},
		{"Test105", args{"105"}, 105, true},
		{"Test106", args{"106"}, 106, true},
		{"Test107", args{"107"}, 107, true},
		{"Test108", args{"108"}, 108, true},
		{"Test109", args{"109"}, 109, true},
		{"Test110", args{"110"}, 110, true},
		{"Test111", args{"111"}, 111, true},
		{"Test112", args{"112"}, 112, true},
		{"Test113", args{"113"}, 113, true},
		{"Test114", args{"114"}, 114, true},
		{"Test115", args{"115"}, 115, true},
		{"Test116", args{"116"}, 116, true},
		{"Test117", args{"117"}, 117, true},
		{"Test118", args{"118"}, 118, true},
		{"Test119", args{"119"}, 119, true},
		{"Test120", args{"120"}, 120, true},
		{"Test121", args{"121"}, 121, true},
		{"Test122", args{"122"}, 122, true},
		{"Test123", args{"123"}, 123, true},
		{"Test124", args{"124"}, 124, true},
		{"Test125", args{"125"}, 125, true},
		{"Test126", args{"126"}, 126, true},
		{"Test127", args{"127"}, 127, true},
		{"Test128", args{"128"}, 128, true},
		{"Test129", args{"129"}, 0, false},
		{"Test130", args{"0x5"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := parseIPv6Length(tt.args.length)
			assert.Equalf(t, tt.want, got, "parseIPv6Length(%v)", tt.args.length)
			assert.Equalf(t, tt.wantOk, gotOk, "parseIPv6Length(%v)", tt.args.length)

		})
	}
}

func Test_v6SplitLen(t *testing.T) {
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
				"2600:c3f:9a2:0::/64",
			},
			[2]string{"2600:c3f:9a2:0::", "64"},
			true,
		},
		{
			"Test1",
			args{
				"2600:c3f:9a2:0::64",
			},
			[2]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := v6SplitLen(tt.args.pfx)
			assert.Equalf(t, tt.want, got, "v6SplitLen(%v)", tt.args.pfx)
			assert.Equalf(t, tt.wantOk, got1, "v6SplitLen(%v)", tt.args.pfx)
		})
	}
}
