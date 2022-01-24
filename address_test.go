package ipnet

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAddress_Bytes(t *testing.T) {
	tests := []struct {
		name string
		addr address
		want u16bytes
	}{
		{
			name: "Test",
			addr: address{0x0, 0xa0a0a0a},
			want: u16bytes{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.addr.bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("u16bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_address_Gt(t *testing.T) {
	tests := []struct {
		name string
		addr address
		in   address
		want bool
	}{
		{
			name: "gt-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "gt-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "lt-gt",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "lt-lt",
			addr: address{0xfffffffffffffffe, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "Eq-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "Eq-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.gt(tt.in), "gt(%v)", tt.in)
		})
	}
}

func Test_address_Ge(t *testing.T) {
	tests := []struct {
		name string
		addr address
		in   address
		want bool
	}{
		{
			name: "gt-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "gt-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "gt-Eq",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "lt-gt",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "lt-lt",
			addr: address{0xfffffffffffffffe, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "lt-Eq",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "Eq-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "Eq-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "Eq-Eq",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.ge(tt.in), "gt(%v)", tt.in)
		})
	}
}

func Test_address_Lt(t *testing.T) {
	tests := []struct {
		name string
		addr address
		in   address
		want bool
	}{
		{
			name: "gt-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "gt-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "lt-gt",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "lt-lt",
			addr: address{0xfffffffffffffffe, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "Eq-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "Eq-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.lt(tt.in), "gt(%v)", tt.in)
		})
	}
}

func Test_address_Le(t *testing.T) {
	tests := []struct {
		name string
		addr address
		in   address
		want bool
	}{
		{
			name: "gt-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "gt-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "gt-Eq",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			want: false,
		},
		{
			name: "lt-gt",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: true,
		},
		{
			name: "lt-lt",
			addr: address{0xfffffffffffffffe, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "lt-Eq",
			addr: address{0xfffffffffffffffe, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "Eq-gt",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xfffffffffffffffe},
			want: false,
		},
		{
			name: "Eq-lt",
			addr: address{0xffffffffffffffff, 0xfffffffffffffffe},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
		{
			name: "Eq-Eq",
			addr: address{0xffffffffffffffff, 0xffffffffffffffff},
			in:   address{0xffffffffffffffff, 0xffffffffffffffff},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.le(tt.in), "gt(%v)", tt.in)
		})
	}
}

func Test_address_Or(t *testing.T) {
	tests := []struct {
		name  string
		addr  address
		input address
		want  address
	}{
		{
			name:  "Or1",
			addr:  address{0xffffffff00000000, 0xffffffff00000000},
			input: address{0xffffffff, 0xffffffff},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
		{
			name:  "Or2",
			addr:  address{0xffffffff, 0xffffffff},
			input: address{0xffffffff00000000, 0xffffffff00000000},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
		{
			name:  "Or3",
			addr:  address{0xf0f0f0f0f0f0f0f0, 0xf0f0f0f0f0f0f0f0},
			input: address{0x0f0f0f0f0f0f0f0f, 0x0f0f0f0f0f0f0f0f},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.or(tt.input), "or(%v)", tt.input)
		})
	}
}

func Test_address_Xor(t *testing.T) {
	tests := []struct {
		name  string
		addr  address
		input address
		want  address
	}{
		{
			name:  "Xor1",
			addr:  address{0xffffffff00000000, 0xffffffff00000000},
			input: address{0xffffffff, 0xffffffff},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
		{
			name:  "Xor2",
			addr:  address{0xffffffff, 0xffffffff},
			input: address{0xffffffff00000000, 0xffffffff00000000},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
		{
			name:  "Xor3",
			addr:  address{0xf0f0f0f0f0f0f0f0, 0xf0f0f0f0f0f0f0f0},
			input: address{0x0f0f0f0f0f0f0f0f, 0x0f0f0f0f0f0f0f0f},
			want:  address{0xffffffffffffffff, 0xffffffffffffffff},
		},
		{
			name:  "Xor4",
			addr:  address{0xffffffffffffffff, 0xffffffffffffffff},
			input: address{0xfe, 0xfe},
			want:  address{0xffffffffffffff01, 0xffffffffffffff01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.addr.xor(tt.input), "or(%v)", tt.input)
		})
	}
}

func Test_address_add(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name   string
		addr   address
		args   args
		want   address
		wantOk bool
	}{
		{
			name:   "Standard",
			addr:   address{0, 0},
			args:   args{i: 1},
			want:   address{0, 1},
			wantOk: true,
		},
		{
			name:   "OverflowLower",
			addr:   address{0, 0xffffffffffffffff},
			args:   args{i: 1},
			want:   address{1, 0},
			wantOk: true,
		},
		{
			name:   "OverflowLowerHigher",
			addr:   address{0xffffffffffffffff, 0xffffffffffffffff},
			args:   args{i: 1},
			want:   address{0xffffffffffffffff, 0xffffffffffffffff},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := tt.addr.add(tt.args.i)
			assert.Equalf(t, tt.want, got, "add(%v)", tt.args.i)
			assert.Equalf(t, tt.wantOk, gotOk, "add(%v)", tt.args.i)
		})
	}
}

func Test_address_subtract(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name   string
		addr   address
		args   args
		want   address
		wantOk bool
	}{
		{
			name:   "Standard",
			addr:   address{0, 1},
			args:   args{i: 1},
			want:   address{0, 0},
			wantOk: true,
		},
		{
			name:   "UnderflowLower",
			addr:   address{0xffffffffffffffff, 0},
			args:   args{i: 1},
			want:   address{0xfffffffffffffffe, 0xffffffffffffffff},
			wantOk: true,
		},
		{
			name:   "UnderflowLowerHigher",
			addr:   address{0, 0},
			args:   args{i: 1},
			want:   address{0, 0},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := tt.addr.subtract(tt.args.i)
			assert.Equalf(t, tt.want, got, "subtract(%v)", tt.args.i)
			assert.Equalf(t, tt.wantOk, gotOk, "subtract(%v)", tt.args.i)
		})
	}
}
