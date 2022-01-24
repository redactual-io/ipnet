package ipnet

type orders int

const (
	higher orders = iota //0
	lower                //1
)

type u16bytes [16]uint8

type address [2]uint64

func (addr address) add(i uint64) (address, bool) {
	if (addr[lower] + i) < addr[lower] { // lower overflow
		if addr[higher] == 0xffffffffffffffff { // higher overflow
			return addr, false
		}
		addr[higher]++
	}
	addr[lower] = addr[lower] + i
	return addr, true
}

func (addr address) subtract(i uint64) (address, bool) {
	if (addr[lower] - i) > addr[lower] { // lower underflow
		if addr[higher] == 0x0 { // higher underflow
			return addr, false
		}
		addr[higher]--
	}
	addr[lower] = addr[lower] - i
	return addr, true
}

func (addr *address) set(h, l uint64) {
	addr[higher], addr[lower] = h, l
}

func (addr address) or(input address) address {
	addr[higher], addr[lower] = addr[higher]|input[higher], addr[lower]|input[lower]
	return addr
}

func (addr address) xor(input address) address {
	addr[higher], addr[lower] = addr[higher]^input[higher], addr[lower]^input[lower]
	return addr
}

func (addr address) bytes() u16bytes {
	const m = 0xffffffffffffff00
	var shifts = [8]int{56, 48, 40, 32, 24, 16, 8, 0}
	var b u16bytes
	for set, order := range []orders{higher, lower} {
		for idx, shift := range shifts {
			b[idx+(set*8)] = uint8((addr[order]>>shift | m) ^ m)
		}
	}
	return b
}

func (addr address) gt(in address) bool {
	if addr[higher] > in[higher] {
		return true
	}
	if addr[higher] == in[higher] && addr[lower] > in[lower] {
		return true
	}
	return false
}

func (addr address) ge(in address) bool {
	if addr[higher] > in[higher] {
		return true
	}
	if addr[higher] == in[higher] && addr[lower] >= in[lower] {
		return true
	}
	return false
}

func (addr address) lt(in address) bool {
	if addr[higher] < in[higher] {
		return true
	}
	if addr[higher] == in[higher] && addr[lower] < in[lower] {
		return true
	}
	return false
}

func (addr address) le(in address) bool {
	if addr[higher] < in[higher] {
		return true
	}
	if addr[higher] == in[higher] && addr[lower] <= in[lower] {
		return true
	}
	return false
}
